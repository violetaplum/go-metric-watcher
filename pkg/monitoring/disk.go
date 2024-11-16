package monitoring

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/disk"
	"time"
)

type DiskMetric struct {
	// 디스크 사용량 관련 메트릭
	Path        string  `json:"path"`
	Total       uint64  `json:"total"`        // 전체 디스크 용량 (bytes)
	Used        uint64  `json:"used"`         // 사용 중인 용량 (bytes)
	Free        uint64  `json:"free"`         // 사용 가능한 용량 (bytes)
	UsedPercent float64 `json:"used_percent"` // 사용률 (%)

	// 디스크 I/O 관련 메트릭
	ReadCount  uint64 `json:"read_count"`  // 읽기 작업 횟수
	WriteCount uint64 `json:"write_count"` // 쓰기 작업 횟수
	ReadBytes  uint64 `json:"read_bytes"`  // 읽은 바이트 수
	WriteBytes uint64 `json:"write_bytes"` // 쓴 바이트 수
	ReadTime   uint64 `json:"read_time"`   // 읽기 작업에 걸린 시간 (ms)
	WriteTime  uint64 `json:"write_time"`  // 쓰기 작업에 걸린 시간 (ms)
	IOTime     uint64 `json:"io_time"`     // I/O 작업에 걸린 총 시간 (ms)
	WeightedIO uint64 `json:"weighted_io"` // 가중치가 부여된 I/O 시간
}

type DiskMonitor struct {
	path            string
	lastMeasurement time.Time
}

func NewDiskMonitor(path string) *DiskMonitor {
	return &DiskMonitor{
		path: path,
	}
}

func (d *DiskMonitor) Collect() (*DiskMetric, error) {
	diskStat, err := disk.Usage(d.path)
	if err != nil {
		return nil, fmt.Errorf("failed to collect disk metrics for %s: %w", d.path, err)
	}

	ioStats, err := disk.IOCounters()
	if err != nil {
		return nil, fmt.Errorf("failed to collect disk IO metrics: %w", err)
	}

	d.lastMeasurement = time.Now()

	fmt.Println(diskStat, ioStats)

	//todo: 수집 메트릭을 프로메테우스에 저장하는 로직 추가

	return NewDiskMetric(diskStat, ioStats), nil
}

func NewDiskMetric(diskStat *disk.UsageStat, ioStats map[string]disk.IOCountersStat) *DiskMetric {
	metric := &DiskMetric{
		Path:        diskStat.Path,
		Total:       diskStat.Total,
		Used:        diskStat.Free,
		Free:        diskStat.Free,
		UsedPercent: diskStat.UsedPercent,
	}

	for _, ioStat := range ioStats {
		metric.ReadCount = ioStat.ReadCount
		metric.WriteCount = ioStat.WriteCount
		metric.ReadBytes = ioStat.ReadBytes
		metric.WriteBytes = ioStat.WriteBytes
		metric.ReadTime = ioStat.ReadTime
		metric.WriteTime = ioStat.WriteTime
		metric.IOTime = ioStat.IoTime
		metric.WeightedIO = ioStat.WeightedIO
		break
	}

	return metric
}
