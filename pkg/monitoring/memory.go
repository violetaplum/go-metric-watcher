package monitoring

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/mem"
	"time"
)

type MemoryMonitor struct {
	lastMeasurement time.Time
}

func NewMemoryMonitor() *MemoryMonitor {
	return &MemoryMonitor{}
}

type MemoryMetrics struct {
	// 기본 메모리 정보
	Total       uint64  `json:"total"`        // 전체 메모리 용량 (bytes)
	Available   uint64  `json:"available"`    // 사용 가능한 메모리 (bytes)
	Used        uint64  `json:"used"`         // 사용 중인 메모리 (bytes)
	Free        uint64  `json:"free"`         // 여유 메모리 (bytes)
	UsedPercent float64 `json:"used_percent"` // 메모리 사용률 (%)

	// 상세 메모리 정보
	Active     uint64 `json:"active"`      // 활성 상태의 메모리
	Inactive   uint64 `json:"inactive"`    // 비활성 상태의 메모리
	Wired      uint64 `json:"wired"`       // Wired 메모리 (macOS)
	Cached     uint64 `json:"cached"`      // 캐시된 메모리
	BufferSize uint64 `json:"buffer_size"` // 버퍼 크기
}

func NewMemoryMetric(vmStat *mem.VirtualMemoryStat) *MemoryMetrics {
	return &MemoryMetrics{
		Total:       vmStat.Total,
		Available:   vmStat.Available,
		Used:        vmStat.Used,
		Free:        vmStat.Free,
		UsedPercent: vmStat.UsedPercent,
		Active:      vmStat.Active,
		Inactive:    vmStat.Inactive,
		Wired:       vmStat.Wired,
		Cached:      vmStat.Cached,
		BufferSize:  vmStat.Buffers,
	}
}

func (m *MemoryMonitor) Collect() (*MemoryMetrics, error) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, fmt.Errorf("failed to collect memory metris: %w", err)
	}

	m.lastMeasurement = time.Now()
	fmt.Println(vmStat)

	return NewMemoryMetric(vmStat), nil
}
