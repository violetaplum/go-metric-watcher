package monitoring

import (
	"github.com/shirou/gopsutil/cpu"
	"time"
)

type CPUCollector struct{}

func NewCPUCollector() *CPUCollector {
	return &CPUCollector{}
}

func (c *CPUCollector) Connect() (float64, error) {
	percentages, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0, err
	}
	return percentages[0], err
}
