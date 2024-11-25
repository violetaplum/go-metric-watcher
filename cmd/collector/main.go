package main

import (
	"github.com/violetaplum/go-metric-watcher/internal/service"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 매트릭 프로세서 생성 (15초 간격으로 수집)
	processor := service.NewMetricProcessor(15 * time.Second)

	// 종료 채널 생성
	stopCh := make(chan struct{})

	// Prometheus 메트릭 서버 시작 (2112 포트)
	processor.StartCollect(stopCh)
	sigCh := make(chan os.Signal, 1)

	// 종료 시그널 처리
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("Metric collector started. Prometheus endpoint: http://localhost:2112/metrics")

	// 종료 시그널 대기
	<-sigCh
	log.Println("Shutting down collector...")

	// 종료 처리
	close(stopCh)

}
