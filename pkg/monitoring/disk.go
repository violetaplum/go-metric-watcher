package monitoring

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/disk"
	"time"
)

type DiskMonitor struct {
	path            string
	lastMeasurement time.Time
}

func NewDiskMonitor(path string) *DiskMonitor {
	return &DiskMonitor{
		path: path,
	}
}

func (d *DiskMonitor) Collect() error {
	diskStat, err := disk.Usage(d.path)
	if err != nil {
		return fmt.Errorf("failed to collect disk metrics for %s: %w", d.path, err)
	}

	ioStats, err := disk.IOCounters()
	if err != nil {
		return fmt.Errorf("failed to collect disk IO metrics: %w", err)
	}

	d.lastMeasurement = time.Now()

	fmt.Println(diskStat, ioStats)

	//todo: 수집 메트릭을 프로메테우스에 저장하는 로직 추가

	return nil
}
