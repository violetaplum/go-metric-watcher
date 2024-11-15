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

func (m *MemoryMonitor) Collect() error {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return fmt.Errorf("failed to collect memory metris: %w", err)
	}

	m.lastMeasurement = time.Now()
	//todo: 수집 메트릭을 프로메테우스에 저장하는 로직 추가
	fmt.Println(vmStat)

	return nil
}
