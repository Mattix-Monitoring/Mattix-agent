package disk

import (
	"context"

	"github.com/shirou/gopsutil/v4/disk"
)

func NewCollector() *Collector {
	return &Collector{}
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
	skip := map[string]struct{}{
		"tmpfs":    {},
		"devtmpfs": {},
		"proc":     {},
		"sysfs":    {},
		"overlay":  {},
		"squashfs": {},
		"cgroup":   {},
		"cgroup2":  {},
	}

	for _, partition := range partitions {
		usage, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			continue
		}
		if _, ok := skip[partition.Fstype]; ok {
			continue
		}

		disks = append(disks, DiskData{
			MountPoint: partition.Mountpoint,
			Total:      usage.Total,
			Used:       usage.Used,
			Free:       usage.Free,
		})
	}

	return disks, nil
}
