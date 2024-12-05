package model

import (
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

	// 네트워크 관련 (기본적인 송수신 정보만)
	NetworkBytesRecv uint64 // 수신된 총 바이트
	NetworkBytesSent uint64 // 송신된 총 바이트
}

type SystemMetricAverage struct {
	CPUUsage    float64
	MemoryUsage float64
	DiskUsage   float64
}
