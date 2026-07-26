package http

import "time"

type APIResponse struct {
	System   SystemInfo     `json:"system"`
	Hardware map[string]any `json:"hardware"`
}

type SystemInfo struct {
	Hostname  string    `json:"hostname"`
	Uptime    uint64    `json:"uptime"`
	UpdatedAt time.Time `json:"updated_at"`
}
