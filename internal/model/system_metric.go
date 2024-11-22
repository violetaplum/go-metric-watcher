package model

import (
	watcherPb "github.com/violetaplum/go-metric-watcher/proto/gen/go/metrics/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type SystemMetric struct {
	Timestamp time.Time // 메트릭이 수집된 시간

	// CPU 관련
	CPUUsage float64 // CPU 사용률 (%)

	// 메모리 관련
	MemoryUsage float64 // 메모리 사용률 (%)
	MemoryTotal uint64  // 전체 메모리 크기 (bytes)
	MemoryFree  uint64  // 사용 가능한 메모리 (bytes)

	// 디스크 관련
	DiskUsage float64 // 디스크 사용률 (%)
	DiskTotal uint64  // 전체 디스크 크기 (bytes)
	DiskFree  uint64  // 사용 가능한 디스크 공간 (bytes)
}

func (m *SystemMetric) ToProto() *watcherPb.SystemMetric {
	return &watcherPb.SystemMetric{
		Timestamp:   timestamppb.New(m.Timestamp),
		CpuUsage:    m.CPUUsage,
		MemoryUsage: m.MemoryUsage,
		MemoryTotal: m.MemoryTotal,
		MemoryFree:  m.MemoryFree,
		DiskUsage:   m.DiskUsage,
		DiskTotal:   m.DiskTotal,
		DiskFree:    m.DiskFree,
	}
}
