package monitoring

import (
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_CPU(t *testing.T) {
	cpuMonitor := NewCPUMonitor()
	metrics, err := cpuMonitor.Collect()
	require.NoError(t, err)
	t.Log("cores :: ", metrics.Cores)
	t.Log("usage :: ", metrics.Usage)
}

func Test_Disk(t *testing.T) {
	partitions, err := disk.Partitions(true)
	require.NoError(t, err)

	for _, partition := range partitions {
		t.Log("Device: ", partition.Device)
		t.Log("Mount point: ", partition.Mountpoint)
		t.Log("File system type: ", partition.Fstype)
		t.Log("Options: ", partition.Opts)

	}

	diskMonitor := NewDiskMonitor("/Users/aspyn")
	diskData, err := diskMonitor.Collect()
	require.NoError(t, err)

	t.Log("disk data: ", diskData)

}

func Test_Memory(t *testing.T) {
	memMonitor := NewMemoryMonitor()
	mem, err := memMonitor.Collect()
	require.NoError(t, err)

	t.Log("Active: ", mem.Active)
	t.Log("Free: ", mem.Free)
	t.Log("BufferSize: ", mem.BufferSize)
	t.Log("Available: ", mem.Available)
	t.Log("Used: ", mem.Used)
	t.Log("Cached: ", mem.Cached)
	t.Log("Inactive: ", mem.Inactive)
	t.Log("Total: ", mem.Total)
	t.Log("UsedPercent: ", mem.UsedPercent)
	t.Log("Wired: ", mem.Wired)
}
