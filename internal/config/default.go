package config

import (
	"time"
)

func Default() Config {
	return Config{
		Server: struct {
			Address string "yaml:\"address\""
			Port    int    "yaml:\"port\""
		}{
			Address: "0.0.0.0",
			Port:    8080,
		},
		Interval: struct {
			Fast time.Duration "yaml:\"fast\""
			Slow time.Duration "yaml:\"slow\""
		}{
			Fast: time.Second,
			Slow: time.Minute,
		},
		Disk: struct {
			IgnoreFS []string "yaml:\"ignore_fs\""
		}{
			IgnoreFS: []string{"tmpfs", "devtmpfs", "proc", "sysfs", "overlay", "squashfs", "cgroup", "cgroup2"},
		},
		Network: struct {
			IgnoreInterfaces []string "yaml:\"ignore_interfaces\""
		}{
			IgnoreInterfaces: []string{"lo", "docker", "br-", "veth"},
		},
	}
}
