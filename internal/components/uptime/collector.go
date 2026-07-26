package uptime

import (
	"bufio"
	"context"
	"os"
	"strconv"
	"strings"
)

func NewCollector() *Collector {
	return &Collector{}
}

func (c *Collector) Name() string {
	return "uptime"
}

func (c *Collector) Collect(ctx context.Context) (any, error) {
	file, err := os.Open("/proc/uptime")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if !scanner.Scan() {
		return 0, scanner.Err()
	}

	fields := strings.Fields(scanner.Text())

	uptime, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		return 0, err
	}

	return uint64(uptime), nil
}
