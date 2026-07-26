package cpu

import (
	"context"
	"runtime"
	"strconv"
	"strings"
)

func NewCollector() *Collector {
	numCores := runtime.NumCPU()
	return &Collector{
		cores: make([]coreState, numCores),
	}
}

func (c *Collector) Name() string {
	return "cpu"
}

func (c *Collector) Collect(ctx context.Context) (any, error) {
	stats, err := ParseCPUStat("/proc/stat")
	if err != nil {
		return nil, err
	}

	numCores := runtime.NumCPU()

	data := CPUData{
		Cores: make([]Core, numCores),
	}

	for i := range numCores {
		data.Cores[i].ID = i
	}

	for _, stat := range stats {
		if stat.Name == "cpu" {
			data.Usage = CalculateUsage(stat, &c.prevTotal, &c.prevIdle)
			continue
		}
		idStr := strings.TrimPrefix(stat.Name, "cpu")
		id, err := strconv.Atoi(idStr)
		if err != nil || id >= len(c.cores) {
			continue
		}
		data.Cores[id].Usage = CalculateUsage(stat, &c.cores[id].prevTotal, &c.cores[id].prevIdle)
	}
	return data, nil
}
