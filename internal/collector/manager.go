package collector

import (
	"context"
	"log"
	"sync"
	"time"
)

type Component interface {
	Name() string
	Collect(ctx context.Context) (any, error)
}

type Manager struct {
	mu sync.RWMutex

	metrics        map[string]any
	fastComponents []Component
	slowComponents []Component
}

func NewManager(fast []Component, slow []Component) *Manager {
	return &Manager{
		metrics:        make(map[string]any),
		fastComponents: fast,
		slowComponents: slow,
	}
}

func (m *Manager) Start(ctx context.Context) {
	m.collectBatch(ctx, m.fastComponents)
	m.collectBatch(ctx, m.slowComponents)

	go m.runLoop(ctx, m.fastComponents, time.Second)
	go m.runLoop(ctx, m.slowComponents, time.Minute)
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
