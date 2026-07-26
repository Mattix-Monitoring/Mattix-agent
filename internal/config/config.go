package config

import "time"

type Config struct {
	Server struct {
		Address string `yaml:"address"`
		Port    int    `yaml:"port"`
	} `yaml:"server"`

	Interval struct {
		Fast time.Duration `yaml:"fast"`
		Slow time.Duration `yaml:"slow"`
	} `yaml:"interval"`

	Disk struct {
		IgnoreFS []string `yaml:"ignore_fs"`
	} `yaml:"disk"`

	Network struct {
		IgnoreInterfaces []string `yaml:"ignore_interfaces"`
	} `yaml:"network"`
}
