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

	}

	//diskMonitor := NewDiskMonitor("")

}
