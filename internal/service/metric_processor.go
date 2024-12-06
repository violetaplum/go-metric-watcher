package service

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/violetaplum/go-metric-watcher/domain"
	"github.com/violetaplum/go-metric-watcher/internal/model"
	"github.com/violetaplum/go-metric-watcher/internal/repository"
	"github.com/violetaplum/go-metric-watcher/pkg/monitoring"
	"github.com/violetaplum/go-metric-watcher/pkg/notifier"
	"gorm.io/gorm"
	"log"
	"net/http"
	"sync"
	"time"
)

type MetricProcessor struct {
	mu              sync.RWMutex
	cpuMonitor      *monitoring.CPUMonitor
	memoryMonitor   *monitoring.MemoryMonitor
	diskMonitor     *monitoring.DiskMonitor
	networkMonitor  *monitoring.NetworkMonitor
	metrics         []model.SystemMetric
	collectionTime  time.Duration
	promDB          *repository.PrometheusDB
	alertService    *notifier.AlertService
	alertRepository domain.AlertHistoryRepository
}

func NewMetricProcessor(collectionTime time.Duration, db *gorm.DB) *MetricProcessor {
	return &MetricProcessor{
		cpuMonitor:      monitoring.NewCPUMonitor(),
		memoryMonitor:   monitoring.NewMemoryMonitor(),
		diskMonitor:     monitoring.NewDiskMonitor("/"),
		collectionTime:  collectionTime,
		promDB:          repository.NewPrometheusDB(),
		networkMonitor:  monitoring.NewNetworkMonitor(),
		alertService:    notifier.NewAlertService(model.DefaultConfig()),
		alertRepository: repository.NewAlertHistoryRepository(db),
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

	networkMetrics, err := mp.networkMonitor.Collect()
	if err != nil {
		return err
	}

	mp.promDB.SaveNetworkMetrics(networkMetrics)

	// 모든 인터페이스의 송수신 합계 계산
	var totalBytesRecv uint64
	var totalBytesSent uint64
	for _, metric := range networkMetrics {
		totalBytesRecv += metric.BytesRecv
		totalBytesSent += metric.BytesSent
	}

	metric := model.SystemMetric{
		Timestamp:        time.Now(),
		CPUUsage:         cpuMetrics.Usage,
		MemoryUsage:      cpuMetrics.Usage,
		MemoryTotal:      memoryMetrics.Total,
		MemoryFree:       memoryMetrics.Free,
		DiskUsage:        diskMetrics.UsedPercent,
		DiskTotal:        diskMetrics.Total,
		DiskFree:         diskMetrics.Free,
		NetworkBytesRecv: totalBytesRecv,
		NetworkBytesSent: totalBytesSent,
	}

	mp.metrics = append(mp.metrics, metric)

	if err := mp.alertService.CheckMetricsAndAlert(metric); err != nil {
		log.Printf("Failed to check alerts: %v", err)
	}

	// 알림 히스토리 저장 함수 추가
	if err := mp.checkAndSaveAlerts(metric); err != nil {
		log.Printf("Failed to save alert history: %v", err)
	}

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

func (mp *MetricProcessor) checkAndSaveAlerts(metric model.SystemMetric) error {
	ctx := context.Background()
	config := model.DefaultConfig()

	if metric.CPUUsage > config.Thresholds.CPUUsage {
		alertHistory := &model.AlertHistory{
			Time:           metric.Timestamp,
			AlertRuleID:    1,
			MetricName:     "cpu",
			ThresholdValue: config.Thresholds.CPUUsage,
			Status:         "triggered",
			Description:    fmt.Sprintf("CPU usage (%.2f%%) exceeded threshold (%.2f%%)", metric.CPUUsage, config.Thresholds.CPUUsage),
			TargetSystem:   "system",
			Severity:       "critical",
		}
		if err := mp.alertRepository.SaveAlert(ctx, alertHistory); err != nil {
			log.Printf("Failed to save cpu alert: %v", err)
		}
	}

	if metric.MemoryUsage > config.Thresholds.MemoryUsage {
		alertHistory := &model.AlertHistory{
			Time:           metric.Timestamp,
			AlertRuleID:    2,
			MetricName:     "memory",
			ThresholdValue: config.Thresholds.MemoryUsage,
			Status:         "triggered",
			Description: fmt.Sprintf("Memory usage (%.2f%%) exceeded threshold (%.2f%%)",
				metric.MemoryUsage, config.Thresholds.MemoryUsage),
			TargetSystem: "system",
			Severity:     "warning",
		}
		if err := mp.alertRepository.SaveAlert(ctx, alertHistory); err != nil {
			log.Printf("Failed to save memory alert: %v", err)
		}
	}

	if metric.DiskUsage > config.Thresholds.DiskUsage {
		alertHistory := &model.AlertHistory{
			Time:           metric.Timestamp,
			AlertRuleID:    3, // Disk Rule ID
			MetricName:     "disk",
			MetricValue:    metric.DiskUsage,
			ThresholdValue: config.Thresholds.DiskUsage,
			Status:         "triggered",
			Description: fmt.Sprintf("Disk usage (%.2f%%) exceeded threshold (%.2f%%)",
				metric.DiskUsage, config.Thresholds.DiskUsage),
			TargetSystem: "system",
			Severity:     "warning",
		}
		if err := mp.alertRepository.SaveAlert(ctx, alertHistory); err != nil {
			log.Printf("Failed to save disk alert: %v", err)
		}
	}

	unresolvedAlerts, err := mp.alertRepository.GetUnresolvedAlerts(ctx)
	if err != nil {
		return fmt.Errorf("failed to get unresolved alerts: %v", err)
	}

	for _, alert := range unresolvedAlerts {
		isResolved := false
		switch alert.MetricName {
		case "cpu":
			isResolved = metric.CPUUsage < alert.ThresholdValue
		case "memory":
			isResolved = metric.MemoryUsage < alert.ThresholdValue
		case "disk":
			isResolved = metric.DiskUsage < alert.ThresholdValue
		}
		if isResolved {
			now := time.Now()
			alert.Status = "resolved"
			alert.ResolvedAt = &now
			if err := mp.alertRepository.UpdateAlert(ctx, &alert); err != nil {
				log.Printf("Failed to update resolved alert: %v", err)
			}
		}
	}

	return nil
}
