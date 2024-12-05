package service

import (
	"context"
	"testing"
	"time"
)

//# 모든 벤치마크 실행
//go test -bench=. -benchmem ./internal/service
//
//# 메모리 프로파일링
//go test -bench=. -benchmem -memprofile=mem.prof ./internal/service
//
//# CPU 프로파일링
//go test -bench=. -benchmem -cpuprofile=cpu.prof ./internal/service
//
//# 특정 벤치마크만 실행
//go test -bench=BenchmarkMetricProcessor_FullSystem -benchmem ./internal/service

func BenchmarkMetricProcessor_Collect(b *testing.B) {
	tests := []struct {
		name           string
		collectionTime time.Duration
	}{
		{"FastCollection_100ms", 100 * time.Millisecond},
		{"NormalCollection_1s", 1 * time.Second},
		{"SlowCollection_5s", 5 * time.Second},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			processor := NewMetricProcessor(tt.collectionTime)

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				err := processor.collect()
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func BenchmarkMetricProcessor_ConcurrentAccess(b *testing.B) {
	processor := NewMetricProcessor(time.Second)
	stopCh := make(chan struct{})

	go processor.StartCollect(stopCh)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = processor.GetLatestMetric()
			_ = processor.GetAverages()
		}
	})

	close(stopCh)
}

func BenchmarkMEtricProcessor_FullSystem(b *testing.B) {
	processor := NewMetricProcessor(100 * time.Millisecond)
	stopCh := make(chan struct{})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go processor.StartCollect(stopCh)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		select {
		case <-ctx.Done():
			b.Fatal("timeout")
		default:
			_, _ = processor.GetLatestMetric()

			end := time.Now()
			start := end.Add(-1 * time.Hour)
			_ = processor.GetMetricsByTimeRange(start, end)

			_ = processor.GetAverages()
		}
	}
	close(stopCh)
}

func BenchmarkMetricProcessor_MemoryUsage(b *testing.B) {
	processor := NewMetricProcessor(100 * time.Millisecond)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if err := processor.collect(); err != nil {
			b.Fatal(err)
		}
	}
}
