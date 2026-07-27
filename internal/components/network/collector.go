package network

import (
	"context"
	"net"

	"github.com/matesu777/Mattix/internal/config"
	psnet "github.com/shirou/gopsutil/v4/net"
)

func NewCollector(cfg *config.NetworkConfig) *Collector {
	ignore := make(map[string]struct{}, len(cfg.IgnoreInterfaces))

	for _, inter := range cfg.IgnoreInterfaces {
		ignore[inter] = struct{}{}
	}
	return &Collector{
		prev: make(map[string]Counter),
	}
}

func (c *Collector) Name() string {
	return "network"
}

func (c *Collector) Collect(ctx context.Context) (any, error) {
	interfaces, err := psnet.Interfaces()
	if err != nil {
		return nil, err
	}

	counters, err := psnet.IOCounters(true)
	if err != nil {
		return nil, err
	}

	counterMap := make(map[string]psnet.IOCountersStat)
	for _, counter := range counters {
		counterMap[counter.Name] = counter
	}

	var data []NetworkData

	for _, iface := range interfaces {
		counter, ok := counterMap[iface.Name]
		if !ok {
			continue
		}
		if c.ignored(iface.Name) {
			continue
		}

		network := NetworkData{
			Name: iface.Name,
			MAC:  iface.HardwareAddr,
		}

		// IPv4
		for _, addr := range iface.Addrs {
			if addr.Addr == "" {
				continue
			}

			ip, _, err := net.ParseCIDR(addr.Addr)
			if err != nil {
				continue
			}

			if ipv4 := ip.To4(); ipv4 != nil {
				network.IPv4 = ipv4.String()
				break
			}
		}

		network.RxBytes = counter.BytesRecv
		network.TxBytes = counter.BytesSent

		prev := c.prev[iface.Name]

		if prev.Rx != 0 {
			network.RxSpeed = counter.BytesRecv - prev.Rx
		}

		if prev.Tx != 0 {
			network.TxSpeed = counter.BytesSent - prev.Tx
		}

		c.prev[iface.Name] = Counter{
			Rx: counter.BytesRecv,
			Tx: counter.BytesSent,
		}

		data = append(data, network)
	}

	return data, nil
}

func (c *Collector) ignored(inter string) bool {
	_, ok := c.ignoreInterfaces[inter]
	return ok
}
