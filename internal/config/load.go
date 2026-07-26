package config

import (
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"gopkg.in/yaml.v3"
)

func XDGPath() string {
	return filepath.Join(xdg.ConfigHome, "mattix", "config.yaml")
}

func Save(path string, cfg *Config) error {
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func Load() (*Config, error) {
	path := XDGPath()

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return nil, err
	}

	cfg := Default()

	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := Save(path, &cfg); err != nil {
			return nil, err
		}

		return &cfg, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
