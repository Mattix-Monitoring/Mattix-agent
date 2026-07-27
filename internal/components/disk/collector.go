package disk

import (
	"context"

	"github.com/matesu777/Mattix/internal/config"
	"github.com/shirou/gopsutil/v4/disk"
)

func NewCollector(cfg *config.DiskConfig) *Collector {
	ignore := make(map[string]struct{}, len(cfg.IgnoreFS))

	for _, fs := range cfg.IgnoreFS {
		ignore[fs] = struct{}{}
	}

	return &Collector{
		ignoreFS: ignore,
	}
}

func (c *Collector) Name() string {
	return "disk"
}

func (c *Collector) Collect(ctx context.Context) (any, error) {
	partitions, err := disk.Partitions(false)
	if err != nil {
		return nil, err
	}

	var disks []DiskData

	for _, partition := range partitions {
		usage, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			continue
		}
		if c.ignored(partition.Fstype) {
			continue
		}

		disks = append(disks, DiskData{
			Device:      partition.Device,
			MountPoint:  partition.Mountpoint,
			FsType:      partition.Fstype,
			Total:       usage.Total,
			Used:        usage.Used,
			Free:        usage.Free,
			UsedPercent: usage.UsedPercent,
		})
	}

	return disks, nil
}

func (c *Collector) ignored(fs string) bool {
	_, ok := c.ignoreFS[fs]
	return ok
}
