package service

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/violetaplum/go-metric-watcher/domain/mocks"
	"github.com/violetaplum/go-metric-watcher/internal/model"
	"github.com/violetaplum/go-metric-watcher/pkg/monitoring"
	"runtime"
	"testing"
	"time"
)

func Test_MetricProcessorIntegration(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockProcessorService(ctrl)

	t.Run("GetLatestMetric", func(t *testing.T) {
		// 예상되는 반환값 설정
		expectedMetric := model.SystemMetric{
			Timestamp:   time.Now(),
			CPUUsage:    50.0,
			MemoryUsage: 60.0,
			DiskUsage:   70.0,
			MemoryTotal: 1000,
			MemoryFree:  400,
			DiskTotal:   2000,
			DiskFree:    800,
		}

		// mock 동작 설정
		mockService.EXPECT().
			GetLatestMetric().
			Return(expectedMetric, nil)

		// 테스트 실행
		metric, err := mockService.GetLatestMetric()

		// 검증
		assert.NoError(t, err)
		assert.Equal(t, expectedMetric, metric)
	})

	t.Run("GetMetricsByTimeRange", func(t *testing.T) {
		start := time.Now().Add(-1 * time.Hour)
		end := time.Now()

		expectedMetrics := []model.SystemMetric{
			{
				Timestamp:   start.Add(10 * time.Minute),
				CPUUsage:    40.0,
				MemoryUsage: 50.0,
				DiskUsage:   60.0,
			},
			{
				Timestamp:   start.Add(20 * time.Minute),
				CPUUsage:    45.0,
				MemoryUsage: 55.0,
				DiskUsage:   65.0,
			},
		}

		mockService.EXPECT().
			GetMetricsByTimeRange(start, end).
			Return(expectedMetrics)

		metrics := mockService.GetMetricsByTimeRange(start, end)
		assert.Equal(t, expectedMetrics, metrics)

	})

	t.Run("GetAverages", func(t *testing.T) {
		expectedAverages := model.SystemMetricAverage{
			CPUUsage:    45.0,
			MemoryUsage: 55.0,
			DiskUsage:   65.0,
		}

		mockService.EXPECT().
			GetAverages().
			Return(expectedAverages)

		averages := mockService.GetAverages()
		assert.Equal(t, expectedAverages, averages)
	})

}

func Test_MetricProcessorMemoryLeak(t *testing.T) {
	// -short 플래그가 설정 돼 있음연 이 테스트는 건너뜀
	if testing.Short() {
		t.Skip("Skipping memory leak test in short mode")
	}

	processor := NewMetricProcessor(100 * time.Millisecond)
	stopCh := make(chan struct{})

	// 초기 메모리 사용량 기록
	var initialMemory runtime.MemStats
	runtime.ReadMemStats(&initialMemory)

	processor.StartCollect(stopCh)

	// 테스트 시간을 5초로 단축시킨다
	testDuration := 5 * time.Second

	// 일정 시간 동안 실행
	time.Sleep(testDuration)
	close(stopCh)

	// 모든 고루틴이 정리될 시간을 조금 줌
	time.Sleep(100 * time.Millisecond)

	// 최종 메모리 사용량 확인
	var finalMemory runtime.MemStats
	runtime.ReadMemStats(&finalMemory)

	// 메모리 증가가 허용 범위 내인지 확인
	memoryIncrease := finalMemory.Alloc - initialMemory.Alloc
	assert.Less(t, memoryIncrease, uint64(10*1024*1024)) // 10MB 이하로 증가
}

func TestPrometheusRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockPrometheusRepository(ctrl)

	t.Run("SaveMetrics", func(t *testing.T) {
		cpuMetrics := &monitoring.CPUMetrics{
			Usage: 50.0,
			Cores: 4,
		}

		memoryMetrics := &monitoring.MemoryMetric{
			Total:       1000,
			Available:   400,
			Used:        600,
			UsedPercent: 60.0,
		}

		diskMetrics := &monitoring.DiskMetrics{
			Total:       2000,
			Used:        1200,
			Free:        800,
			UsedPercent: 60.0,
		}

		networkMetrics := map[string]*monitoring.NetworkMetric{
			"eth0": {
				Interface:   "eth0",
				BytesSent:   1000,
				BytesRecv:   2000,
				PacketsSent: 100,
				PacketsRecv: 200,
			},
		}

		// 각 Save 함수 호출 예상 설정
		mockRepo.EXPECT().SaveCPUMetrics(cpuMetrics)
		mockRepo.EXPECT().SaveMemoryMetrics(memoryMetrics)
		mockRepo.EXPECT().SaveDiskMetrics(diskMetrics)
		mockRepo.EXPECT().SaveNetworkMetrics(networkMetrics)

		// 메트릭 저장 실행
		mockRepo.SaveCPUMetrics(cpuMetrics)
		mockRepo.SaveMemoryMetrics(memoryMetrics)
		mockRepo.SaveDiskMetrics(diskMetrics)
		mockRepo.SaveNetworkMetrics(networkMetrics)
	})
}
