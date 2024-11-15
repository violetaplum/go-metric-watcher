package monitoring

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"time"
)

type CPUMetrics struct {
	Usage float64
	Cores int
}

type CPUMonitor struct {
	lastMeasurement time.Time
}

func NewCPUMonitor() *CPUMonitor {
	return &CPUMonitor{}
}

func (m *CPUMonitor) Collect() (*CPUMetrics, error) {
	percentage, err := cpu.Percent(time.Second, false)
	if err != nil {
		return nil, fmt.Errorf("failed to get CPU percentage: %w", err)
	}

	cores, err := cpu.Counts(true)
	if err != nil {
		return nil, fmt.Errorf("failed to get CPU cores: %w", err)
	}

	m.lastMeasurement = time.Now()
	return &CPUMetrics{
		Usage: percentage[0],
		Cores: cores,
	}, nil
}

func (m *CPUMonitor) LastMeasurement() time.Time {
	return m.lastMeasurement
}
