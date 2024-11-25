package repository

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/violetaplum/go-metric-watcher/pkg/monitoring"
)

type PrometheusDB struct {
	// cpu 메트릭
	cpuUsage prometheus.Gauge
	cpuCores prometheus.Gauge

	// 메모리 메트릭
	memoryTotal       prometheus.Gauge
	memoryAvailable   prometheus.Gauge
	memoryUsed        prometheus.Gauge
	memoryFree        prometheus.Gauge
	memoryUsedPercent prometheus.Gauge
	memoryActive      prometheus.Gauge
	memoryInactive    prometheus.Gauge
	memoryWired       prometheus.Gauge
	memoryCached      prometheus.Gauge
	memoryBufferSize  prometheus.Gauge

	// 디스크 메트릭
	diskTotal       *prometheus.GaugeVec
	diskUsed        *prometheus.GaugeVec
	diskFree        *prometheus.GaugeVec
	diskUsedPercent *prometheus.GaugeVec
	diskReadCount   *prometheus.GaugeVec
	diskWriteCount  *prometheus.GaugeVec
	diskReadBytes   *prometheus.GaugeVec
	diskWriteBytes  *prometheus.GaugeVec
	diskReadTime    *prometheus.GaugeVec
	diskWriteTime   *prometheus.GaugeVec
	diskIOTime      *prometheus.GaugeVec
	diskWeightedIO  *prometheus.GaugeVec
}

func NewPrometheusDB() *PrometheusDB {
	return &PrometheusDB{
		// cpu 메트릭
		cpuUsage: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "system_cou_usgae_percent",
			Help: "Current CPU usage in percentage",
		}),
		cpuCores: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "system_cpu_cores_total",
			Help: "Total number of CPU cores",
		}),

		// memory 메트릭
		memoryTotal: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "system_memory_total_bytes",
			Help: "Total amount of memory in bytes",
		}),
		memoryAvailable: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "system_memory_available_bytes",
			Help: "Available memory in bytes",
		}),
		memoryUsed: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "system_memory_used_bytes",
			Help: "Used memory in bytes",
		}),
		memoryFree: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "system_memory_free_bytes",
			Help: "Free memory in bytes",
		}),
		memoryUsedPercent: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "system_memory_used_percent",
			Help: "Memory usage percentage",
		}),
		memoryActive: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "system_memory_active_bytes",
			Help: "Active memory in bytes",
		}),
		memoryInactive: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "system_memory_inactive_bytes",
			Help: "Inactive memory in bytes",
		}),
		memoryWired: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "system_memory_wired_bytes",
			Help: "Wired memory in bytes (macOS)",
		}),
		memoryCached: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "system_memory_cached_bytes",
			Help: "Cached memory in bytes",
		}),
		memoryBufferSize: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "system_memory_buffer_bytes",
			Help: "Buffer size in bytes",
		}),

		// disk 메트릭 (레이블 포함)
		diskTotal: promauto.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "system_disk_total_bytes",
				Help: "Total disk space in bytes",
			},
			[]string{"path"}, // path 레이블 모두 추가
		),
		diskUsed: promauto.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "system_disk_used_bytes",
				Help: "Used disk space in bytes",
			},
			[]string{"path"},
		),
		diskFree: promauto.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "system_disk_free_bytes",
				Help: "Free disk space in bytes",
			},
			[]string{"path"},
		),
		diskUsedPercent: promauto.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "system_disk_used_percent",
				Help: "Disk usage percentage",
			},
			[]string{"path"},
		),
		diskReadCount: promauto.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "system_disk_read_count_total",
				Help: "Total number of read operations",
			},
			[]string{"path"},
		),
		diskWriteCount: promauto.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "system_disk_write_count_total",
				Help: "Total number of write operations",
			},
			[]string{"path"},
		),
		diskReadBytes: promauto.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "system_disk_read_bytes_total",
				Help: "Total number of bytes read",
			},
			[]string{"path"},
		),
		diskWriteBytes: promauto.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "system_disk_write_bytes_total",
				Help: "Total number of bytes written",
			},
			[]string{"path"},
		),
		diskReadTime: promauto.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "system_disk_read_time_milliseconds",
				Help: "Time spent reading from disk",
			},
			[]string{"path"},
		),
		diskWriteTime: promauto.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "system_disk_write_time_milliseconds",
				Help: "Time spent writing to disk",
			},
			[]string{"path"},
		),
		diskIOTime: promauto.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "system_disk_io_time_milliseconds",
				Help: "Total time spent on I/O operations",
			},
			[]string{"path"},
		),
		diskWeightedIO: promauto.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "system_disk_weighted_io_milliseconds",
				Help: "Weighted time spent on I/O operations",
			},
			[]string{"path"},
		),
	}
}

func (p *PrometheusDB) SaveCPUMetrics(metrics *monitoring.CPUMetrics) {
	p.cpuUsage.Set(metrics.Usage)
	p.cpuCores.Set(float64(metrics.Cores))
}

func (p *PrometheusDB) SaveMemoryMetrics(metrics *monitoring.MemoryMetric) {
	p.memoryTotal.Set(float64(metrics.Total))
	p.memoryAvailable.Set(float64(metrics.Available))
	p.memoryUsed.Set(float64(metrics.Used))
	p.memoryFree.Set(float64(metrics.Free))
	p.memoryUsedPercent.Set(metrics.UsedPercent)
	p.memoryActive.Set(float64(metrics.Active))
	p.memoryInactive.Set(float64(metrics.Inactive))
	p.memoryWired.Set(float64(metrics.Wired))
	p.memoryCached.Set(float64(metrics.Cached))
	p.memoryBufferSize.Set(float64(metrics.BufferSize))
}

func (p *PrometheusDB) SaveDiskMetrics(metrics *monitoring.DiskMetrics) {
	p.diskTotal.With(prometheus.Labels{
		"path": metrics.Path,
	}).Set(float64(metrics.Total))

	p.diskUsed.With(prometheus.Labels{
		"path": metrics.Path,
	}).Set(float64(metrics.Used))

	p.diskFree.With(prometheus.Labels{
		"path": metrics.Path,
	}).Set(float64(metrics.Free))

	p.diskUsedPercent.With(prometheus.Labels{
		"path": metrics.Path,
	}).Set(metrics.UsedPercent)

	// I/O 메트릭
	p.diskReadCount.With(prometheus.Labels{
		"path": metrics.Path,
	}).Set(float64(metrics.ReadCount))

	p.diskWriteCount.With(prometheus.Labels{
		"path": metrics.Path,
	}).Set(float64(metrics.WriteCount))

	p.diskReadBytes.With(prometheus.Labels{
		"path": metrics.Path,
	}).Set(float64(metrics.ReadBytes))

	p.diskWriteBytes.With(prometheus.Labels{
		"path": metrics.Path,
	}).Set(float64(metrics.WriteBytes))

	p.diskReadTime.With(prometheus.Labels{
		"path": metrics.Path,
	}).Set(float64(metrics.ReadTime))

	p.diskWriteTime.With(prometheus.Labels{
		"path": metrics.Path,
	}).Set(float64(metrics.WriteTime))

	p.diskIOTime.With(prometheus.Labels{
		"path": metrics.Path,
	}).Set(float64(metrics.IOTime))

	p.diskWeightedIO.With(prometheus.Labels{
		"path": metrics.Path,
	}).Set(float64(metrics.WeightedIO))
}
