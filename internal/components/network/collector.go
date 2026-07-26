package network

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func NewCollector() (*Collector, error) {
	iface, err := getDefaultInterface()
	if err != nil {
		return nil, err
	}

	c := &Collector{
		ifaceName: iface.Name,
		mac:       iface.HardwareAddr.String(),
	}

	addrs, err := iface.Addrs()
	if err == nil {
		for _, addr := range addrs {
			if ipNet, ok := addr.(*net.IPNet); ok {
				if ip := ipNet.IP.To4(); ip != nil {
					c.ipv4 = ip.String()
					break
				}
			}
		}
	}

	return c, nil
}

func (c *Collector) Name() string {
	return "network"
}

func getDefaultInterface() (*net.Interface, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp == 0 {
			continue
		}

		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil || len(addrs) == 0 {
			continue
		}

		for _, addr := range addrs {
			if ipNet, ok := addr.(*net.IPNet); ok {
				if ipNet.IP.To4() != nil {
					return &iface, nil
				}
			}
		}
	}

	return nil, fmt.Errorf("nenhuma interface de rede ativa encontrada")
}

func (c *Collector) Collect(ctx context.Context) (any, error) {
	file, err := os.Open("/proc/net/dev")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data := NetworkData{
		Name: c.ifaceName,
		MAC:  c.mac,
		IPv4: c.ipv4,
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, "Inter-") ||
			strings.HasPrefix(line, "face") {
			continue
		}

		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}

		iface := strings.TrimSpace(parts[0])

		if iface != c.ifaceName {
			continue
		}

		fields := strings.Fields(parts[1])

		if len(fields) < 16 {
			return nil, fmt.Errorf("formato inválido em /proc/net/dev")
		}

		rx, err := strconv.ParseUint(fields[0], 10, 64)
		if err != nil {
			return nil, err
		}

		tx, err := strconv.ParseUint(fields[8], 10, 64)
		if err != nil {
			return nil, err
		}

		data.RxBytes = rx
		data.TxBytes = tx

		if c.prevRx != 0 {
			data.RxSpeed = rx - c.prevRx
		}

		if c.prevTx != 0 {
			data.TxSpeed = tx - c.prevTx
		}

		c.prevRx = rx
		c.prevTx = tx

		return data, nil
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("Interface %s not found in /proc/net/dev", c.ifaceName)
}
