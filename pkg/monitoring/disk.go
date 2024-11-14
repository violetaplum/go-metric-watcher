package monitoring

import "github.com/shirou/gopsutil/disk"

type DiskCollector struct{}

func NewDiskCollector() *DiskCollector {
	return &DiskCollector{}
}

func (d *DiskCollector) Collect() (*disk.UsageStat, error) {
	return disk.Usage("/")
}
