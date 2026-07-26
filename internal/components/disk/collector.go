package disk

import (
	"context"
	"golang.org/x/sys/unix"
)

func NewCollector() *Collector {
	return &Collector{}
}

func (c *Collector) Name() string {
	return "disk"
}

func (c *Collector) Collect(ctx context.Context) (any, error) {
	var stat unix.Statfs_t

	err := unix.Statfs("/", &stat)
	if err != nil {
		return nil, err
	}

	total := stat.Blocks * uint64(stat.Bsize)
	free := stat.Bavail * uint64(stat.Bsize)

	data := DiskData{
		Total: total,
		Free:  free,
		Used:  total - free,
	}

	return data, nil
}
