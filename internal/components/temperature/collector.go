package temperature

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func NewCollector() *Collector {
	return &Collector{}
}

func (c *Collector) Name() string {
	return "temperature"
}

func GetSensor() (string, error) {
	entries, err := os.ReadDir("/sys/class/hwmon")
	if err != nil {
		return "", err
	}

	for _, entry := range entries {
		dir := filepath.Join("/sys/class/hwmon", entry.Name())

		nameBytes, err := os.ReadFile(filepath.Join(dir, "name"))
		if err != nil {
			continue
		}

		name := strings.TrimSpace(string(nameBytes))

		if name == "coretemp" || name == "k10temp" {
			return dir, nil
		}
	}

	return "", fmt.Errorf("cpu sensor not found")
}

func (c *Collector) Collect(ctx context.Context) (any, error) {
	sensorDir, err := GetSensor()
	if err != nil {
		return nil, err
	}

	for i := 1; ; i++ {
		labelPath := filepath.Join(sensorDir, fmt.Sprintf("temp%d_label", i))

		labelBytes, err := os.ReadFile(labelPath)
		if err != nil {
			break
		}

		if strings.TrimSpace(string(labelBytes)) != "Package id 0" {
			continue
		}

		inputPath := filepath.Join(sensorDir, fmt.Sprintf("temp%d_input", i))

		tempBytes, err := os.ReadFile(inputPath)
		if err != nil {
			return nil, err
		}

		temp, err := strconv.ParseUint(strings.TrimSpace(string(tempBytes)), 10, 64)
		if err != nil {
			return nil, err
		}

		data := &TemperatureData{
			CpuTemp: temp,
		}
		return data, nil
	}

	return nil, fmt.Errorf("package temperature not found")
}
