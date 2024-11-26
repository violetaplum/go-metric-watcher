package service

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/violetaplum/go-metric-watcher/internal/model"
	"github.com/violetaplum/go-metric-watcher/internal/repository"
	"github.com/violetaplum/go-metric-watcher/pkg/monitoring"
	"log"
	"net/http"
	"sync"
	"time"
)

type MetricProcessor struct {
	mu             sync.RWMutex
	cpuMonitor     *monitoring.CPUMonitor
	memoryMonitor  *monitoring.MemoryMonitor
	diskMonitor    *monitoring.DiskMonitor
	metrics        []model.SystemMetric
	collectionTime time.Duration
	promDB         *repository.PrometheusDB
}

func NewMetricProcessor(collectionTime time.Duration) *MetricProcessor {
	return &MetricProcessor{
		cpuMonitor:     monitoring.NewCPUMonitor(),
		memoryMonitor:  monitoring.NewMemoryMonitor(),
		diskMonitor:    monitoring.NewDiskMonitor("/"),
		collectionTime: collectionTime,
		promDB:         repository.NewPrometheusDB(),
	}
}

// 메트릭 수집 시작
func (mp *MetricProcessor) StartCollect(stopCh <-chan struct{}) {
	// Prometheus HTTP 핸들러 등록
	http.Handle("/metrics", promhttp.Handler())

	// Prometheus 메트릭 서버 시작
	go func() {
		if err := http.ListenAndServe(":2112", nil); err != nil {
			log.Printf("Prometheus metrics server failed: %v ", err)

		}
	}()

	ticker := time.NewTicker(mp.collectionTime)
	defer ticker.Stop()

	for {
		select {
		case <-stopCh:
			return
		case <-ticker.C:
			if err := mp.collect(); err != nil {
				log.Printf("Error collecting metrics: %v", err)

				continue
			}

		}
	}
}

func (mp *MetricProcessor) collect() error {
	cpuMetrics, err := mp.cpuMonitor.Collect()
	if err != nil {
		return err
	}

	mp.promDB.SaveCPUMetrics(cpuMetrics)

	memoryMetrics, err := mp.memoryMonitor.Collect()
	if err != nil {
		return err
	}

	mp.promDB.SaveMemoryMetrics(memoryMetrics)

	diskMetrics, err := mp.diskMonitor.Collect()
	if err != nil {
		return err
	}

	mp.promDB.SaveDiskMetrics(diskMetrics)

	mp.metrics = append(mp.metrics, model.SystemMetric{
		Timestamp:   time.Now(),
		CPUUsage:    cpuMetrics.Usage,
		MemoryUsage: cpuMetrics.Usage,
		MemoryTotal: memoryMetrics.Total,
		MemoryFree:  memoryMetrics.Free,
		DiskUsage:   diskMetrics.UsedPercent,
		DiskTotal:   diskMetrics.Total,
		DiskFree:    diskMetrics.Free,
	})

	return nil
}

// 가장 최근 메트릭 조회
func (mp *MetricProcessor) GetLatestMetric() (model.SystemMetric, error) {
	mp.mu.Lock()
	defer mp.mu.Unlock()

	if len(mp.metrics) == 0 {
		return model.SystemMetric{}, fmt.Errorf("no metrics available")
	}

	return mp.metrics[len(mp.metrics)-1], nil
}

func (mp *MetricProcessor) GetMetricsByTimeRange(start, end time.Time) []model.SystemMetric {
	mp.mu.Lock()
	defer mp.mu.Unlock()

	var result []model.SystemMetric
	for _, m := range mp.metrics {
		if m.Timestamp.After(start) && m.Timestamp.Before(end) {
			result = append(result, m)
		}
	}
	return result
}

// 오래된 메트릭 정리 (24시간 이전 데이터)
//func (mp *MetricProcessor) cleanOldMetrics() {
//	mp.mu.Lock()
//	defer mp.mu.Unlock()
//	cutoff := time.Now().Add(-24 * time.Hour)
//	var newMetrics []model.SystemMetric
//
//	for _, metric := range mp.metrics {
//		if metric.Timestamp.After(cutoff) {
//			newMetrics = append(newMetrics, metric)
//		}
//	}
//	mp.metrics = newMetrics
//}

// 평균 시간 계산
func (mp *MetricProcessor) GetAverages() model.SystemMetricAverage {
	mp.mu.Lock()
	defer mp.mu.Unlock()

	var avg model.SystemMetricAverage
	if len(mp.metrics) == 0 {
		return avg
	}

	for _, m := range mp.metrics {
		avg.CPUUsage += m.CPUUsage
		avg.MemoryUsage += m.MemoryUsage
		avg.DiskUsage += m.DiskUsage
	}

	count := float64(len(mp.metrics))
	avg.CPUUsage /= count
	avg.MemoryUsage /= count
	avg.DiskUsage /= count

	return avg
}
