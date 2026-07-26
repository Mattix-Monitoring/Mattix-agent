package hostname

import (
	"context"
	"os"
)

func NewCollector() *Collector {
	return &Collector{}
}

func (c *Collector) Name() string {
	return "hostname"
}

func (c *Collector) Collect(ctx context.Context) (any, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}
	return hostname, nil
}
