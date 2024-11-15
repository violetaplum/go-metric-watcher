package service

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/violetaplum/go-metric-watcher/internal/model"
	"github.com/violetaplum/go-metric-watcher/pkg/logger"
	"sync"
)

type MetricProcessor struct {
	metrics map[string]prometheus.Gauge
	mu      sync.RWMutex
	logger  *logger.Logger
}

func NewMetricProcessor(logger *logger.Logger) *MetricProcessor {
	return &MetricProcessor{
		metrics: make(map[string]prometheus.Gauge),
		logger:  logger,
	}
}

func (mp *MetricProcessor) Process(ctx context.Context, metric *model.SystemMetric) error {
	mp.mu.Lock()

	return nil
}
