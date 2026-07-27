package collector

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/matesu777/Mattix/internal/config"
)

type Component interface {
	Name() string
	Collect(ctx context.Context) (any, error)
}

type Manager struct {
	mu sync.RWMutex

	config         *config.IntervalConfig
	metrics        map[string]any
	fastComponents []Component
	slowComponents []Component
}

func NewManager(fast []Component, slow []Component, cfg *config.IntervalConfig) *Manager {
	return &Manager{
		config:         cfg,
		metrics:        make(map[string]any),
		fastComponents: fast,
		slowComponents: slow,
	}
}

func (m *Manager) Start(ctx context.Context) {
	m.collectBatch(ctx, m.fastComponents)
	m.collectBatch(ctx, m.slowComponents)

	go m.runLoop(ctx, m.fastComponents, m.config.Fast)
	go m.runLoop(ctx, m.slowComponents, m.config.Slow)
}

func (m *Manager) runLoop(ctx context.Context, comps []Component, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			m.collectBatch(ctx, comps)
		}
	}
}

func (m *Manager) collectBatch(ctx context.Context, comps []Component) {
	for _, c := range comps {
		data, err := c.Collect(ctx)
		if err != nil {
			log.Printf("Error collecting %s: %v\n", c.Name(), err)
			continue
		}
		m.mu.Lock()

		m.metrics[c.Name()] = data
		m.metrics["updated_at"] = time.Now()

		m.mu.Unlock()
	}
}

func (m *Manager) GetMetrics() map[string]any {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.metrics
}
