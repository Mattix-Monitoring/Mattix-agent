package disk

type DiskData struct {
	Device      string  `json:"device"`
	MountPoint  string  `json:"mouint_point"`
	FsType      string  `json:"fsType"`
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"used_percent"`
}

type Collector struct {
	ignoreFS map[string]struct{}
}
