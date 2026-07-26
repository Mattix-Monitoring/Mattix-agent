package temperature

type TemperatureData struct {
	CpuTemp uint64 `json:"cpu"`
}

type Collector struct{}
