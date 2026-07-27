package config

import "time"

type Config struct {
	Server ServerConfig `yaml:"server"`

	Interval IntervalConfig `yaml:"interval"`

	Disk DiskConfig `yaml:"disk"`

	Network NetworkConfig `yaml:"network"`
}

type ServerConfig struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

type IntervalConfig struct {
	Fast time.Duration `yaml:"fast"`
	Slow time.Duration `yaml:"slow"`
}

type DiskConfig struct {
	IgnoreFS []string `yaml:"ignore_fs"`
}

type NetworkConfig struct {
	IgnoreInterfaces []string `yaml:"ignore_interfaces"`
}
