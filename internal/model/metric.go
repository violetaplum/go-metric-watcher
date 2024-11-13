package model

import "time"

type MetricResponse struct {
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
	Value     float64   `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
