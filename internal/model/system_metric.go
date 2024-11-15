package model

import (
	watcherPb "github.com/violetaplum/go-metric-watcher/proto/gen/go/metrics/v1"
	"time"
)

type SystemMetric struct {
	// float32: 32비트 (4바이트)
	// - 정밀도: 약 7자리
	// - 범위: ±1.18E-38 ~ ±3.4E38

	// float64: 64비트 (8바이트)
	// - 정밀도: 약 15자리
	// - 범위: ±2.23E-308 ~ ±1.80E308
	// float64를 선호하는 이유:
	//
	// 더 높은 정밀도 필요
	// 더 넓은 범위의 값 표현 가능
	// Go에서 기본 실수형이 float64
	// 메모리 사용량 차이가 크지 않음
	Type      string
	Value     float64
	Labels    map[string]string
	ServerID  string
	Timestamp time.Time
	Unit      watcherPb.MetricUnit
}

func (m *SystemMetric) ToProto() *watcherPb.SystemMetric {
	return &watcherPb.SystemMetric{
		Type:      m.Type,
		Value:     m.Value,
		Labels:    m.Labels,
		ServerId:  m.ServerID,
		Timestamp: m.Timestamp.Unix(),
		Unit:      m.Unit,
	}
}
