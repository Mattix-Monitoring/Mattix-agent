package components

import (
	"bufio"
	"github.com/matesu777/Mattix/internal/utils"
	"os"
	"strings"
)

type Memory struct {
	Total uint64 `json:"total"`
	Used  uint64 `json:"used"`
	Free  uint64 `json:"free"`
}

func (m *Memory) Scan() error {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "MemTotal:") {
			fields := strings.Fields(line)
			total, err := utils.ConvertToUnit64(fields[1])
			if err != nil {
				return err
			}
			m.Total = total * 1024
		}

		if strings.HasPrefix(line, "MemAvailable:") {
			fields := strings.Fields(line)
			free, err := utils.ConvertToUnit64(fields[1])
			if err != nil {
				return err
			}
			m.Free = free * 1024
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	m.Used = m.Total - m.Free

	return nil
}
