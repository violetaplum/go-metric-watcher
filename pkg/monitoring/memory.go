package monitoring

import "github.com/shirou/gopsutil/mem"

type MemoryCollector struct{}

func NewMemoryCollector() *MemoryCollector {
	return &MemoryCollector{}
}

func (m *MemoryCollector) Collect() (*mem.VirtualMemoryStat, error) {
	return mem.VirtualMemory()
}
