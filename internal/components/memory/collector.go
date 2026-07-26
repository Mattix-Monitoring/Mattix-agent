package memory

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
	return "memory"
}

func (c *Collector) Collect(ctx context.Context) (any, error) {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var total, free uint64
	var foundTotal, foundFree bool

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "MemTotal:") {
			fields := strings.Fields(line)
			val, err := strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				return nil, err
			}
			total = val * 1024
			foundTotal = true

		} else if strings.HasPrefix(line, "MemAvailable:") {
			fields := strings.Fields(line)
			val, err := strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				return nil, err
			}
			free = val * 1024
			foundFree = true

		}

		if foundTotal && foundFree {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	data := MemoryData{
		Total: total,
		Free:  free,
		Used:  total - free,
	}

	return data, nil
}
