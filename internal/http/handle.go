package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/matesu777/Mattix/internal/collector"
)

type Handler struct {
	collector *collector.Manager
}

func ServerNew(c *collector.Manager) *Handler {
	return &Handler{
		collector: c,
	}
}

func (h *Handler) GetMetrics(w http.ResponseWriter, r *http.Request) {
	rawData := h.collector.GetMetrics()

	response := APIResponse{
		System: SystemInfo{
			Hostname:  rawData["hostname"].(string),
			Uptime:    rawData["uptime"].(uint64),
			UpdatedAt: rawData["updated_at"].(time.Time),
		},
		Hardware: make(map[string]any),
	}

	for key, value := range rawData {
		if key != "hostname" && key != "uptime" && key != "updated_at" {
			response.Hardware[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
