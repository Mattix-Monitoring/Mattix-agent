package components

import (
	"time"

	"github.com/matesu777/Mattix/internal/components/cpu"
)

type Metrics struct {
	Hostname string    `json:"hostname"`
	Uptime   uint64    `json:"uptime"`
	UpdateAt time.Time `json:"updateAt"`

	CPU         cpu.Cpu     `json:"cpu"`
	Memory      Memory      `json:"memory"`
	Disk        Disk        `json:"disk"`
	Network     Network     `json:"network"`
	Temperature Temperature `json:"temperature"`
}
