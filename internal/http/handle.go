package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/matesu777/Mattix/internal/collector"
	"github.com/matesu777/Mattix/internal/config"
)

type Handler struct {
	config    *config.ServerConfig
	collector *collector.Manager
}

func ServerNew(cfg *config.ServerConfig, c *collector.Manager) *Handler {
	return &Handler{
		config:    cfg,
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

func (h *Handler) StartServer(cfg *config.ServerConfig, handler http.Handler) error {
	addr := fmt.Sprintf("%s:%d",
		cfg.Address,
		cfg.Port,
	)
	return http.ListenAndServe(addr, handler)
}
