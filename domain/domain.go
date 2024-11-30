package domain

import (
	"github.com/violetaplum/go-metric-watcher/internal/model"
	"github.com/violetaplum/go-metric-watcher/pkg/monitoring"
	"time"
)

//go:generate mockgen -source=domain.go -destination=mocks/mock_domain.go -package=mocks
type PrometheusRepository interface {
	SaveCPUMetrics(metrics *monitoring.CPUMetrics)
	SaveMemoryMetrics(metrics *monitoring.MemoryMetric)
	SaveDiskMetrics(metrics *monitoring.DiskMetrics)
	SaveNetworkMetrics(metrics map[string]*monitoring.NetworkMetric)
}

type ProcessorService interface {
	StartCollect(stopCh <-chan struct{})
	GetLatestMetric() (model.SystemMetric, error)
	GetMetricsByTimeRange(start, end time.Time) []model.SystemMetric
	GetAverages() model.SystemMetricAverage
}

type Notifier interface {
	Send(message string) error
}
