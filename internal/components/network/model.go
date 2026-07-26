package network

type NetworkData struct {
	Name string `json:"name"`
	MAC  string `json:"mac"`
	IPv4 string `json:"ipv4"`

	RxBytes uint64 `json:"rx_bytes"`
	TxBytes uint64 `json:"tx_bytes"`

	RxSpeed uint64 `json:"rx_speed"`
	TxSpeed uint64 `json:"tx_speed"`
}

type Collector struct {
	ifaceName string
	mac       string
	ipv4      string
	prevRx    uint64
	prevTx    uint64
}
