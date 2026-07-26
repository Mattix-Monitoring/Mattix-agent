package cpu

type CPUData struct {
	Usage float64 `json:"usage"`
	Cores []Core  `json:"cores"`
}

type Core struct {
	ID    int     `json:"id"`
	Usage float64 `json:"usage"`
}
type Collector struct {
	prevTotal uint64
	prevIdle  uint64
	cores     []coreState
}

type coreState struct {
	prevTotal uint64
	prevIdle  uint64
}

type CPUStat struct {
	Name      string
	User      uint64
	Nice      uint64
	System    uint64
	Idle      uint64
	Iowait    uint64
	IRQ       uint64
	SoftIRQ   uint64
	Steal     uint64
	prevTotal uint64
	prevIdle  uint64
}
