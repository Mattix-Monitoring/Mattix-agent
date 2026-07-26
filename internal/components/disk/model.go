package disk

type DiskData struct {
	MountPoint string `json:"mouint_point"`
	Total      uint64 `json:"total"`
	Free       uint64 `json:"free"`
	Used       uint64 `json:"used"`
}

type Collector struct{}
