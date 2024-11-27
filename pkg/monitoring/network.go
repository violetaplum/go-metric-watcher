package monitoring

import (
	"fmt"
	"github.com/shirou/gopsutil/net"
	"time"
)

type NetworkMetric struct {
	Interface   string `json:"interface"`
	BytesSent   uint64 `json:"bytes_sent"`
	BytesRecv   uint64 `json:"bytes_recv"`
	PacketsSent uint64 `json:"packets_sent"`
	PacketsRecv uint64 `json:"packets_recv"`
	ErrIn       uint64 `json:"err_in"`
	ErrOut      uint64 `json:"err_out"`
	DropIn      uint64 `json:"drop_in"`
	DropOut     uint64 `json:"drop_out"`
}

type NetworkMonitor struct {
	lastMeasurement time.Time
}

func NewNetworkMonitor() *NetworkMonitor {
	return &NetworkMonitor{}
}

func (m *NetworkMonitor) Collect() (map[string]*NetworkMetric, error) {
	stats, err := net.IOCounters(true) // true for per-interface statistics
	if err != nil {
		return nil, fmt.Errorf("failed to collect network metrics: %v", err)
	}
	// stats 구조체 내용
	//type IOCountersStat struct {
	//	Name        string  // 인터페이스 이름
	//	BytesSent   uint64  // 보낸 바이트 수
	//	BytesRecv   uint64  // 받은 바이트 수
	//	PacketsSent uint64  // 보낸 패킷 수
	//	PacketsRecv uint64  // 받은 패킷 수
	//	Errin       uint64  // 수신 에러 수
	//	Errout      uint64  // 송신 에러 수
	//	Dropin      uint64  // 수신 패킷 드랍 수
	//	Dropout     uint64  // 송신 패킷 드랍 수
	//}

	metrics := make(map[string]*NetworkMetric)

	for _, stat := range stats {
		metrics[stat.Name] = &NetworkMetric{
			Interface:   stat.Name,
			BytesSent:   stat.BytesSent,
			BytesRecv:   stat.BytesRecv,
			PacketsSent: stat.PacketsSent,
			PacketsRecv: stat.PacketsRecv,
			ErrIn:       stat.Errin,
			ErrOut:      stat.Errout,
			DropIn:      stat.Dropin,
			DropOut:     stat.Dropout,
		}
	}

	m.lastMeasurement = time.Now()
	return metrics, nil
}
