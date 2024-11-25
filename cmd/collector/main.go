package main

import (
	"github.com/violetaplum/go-metric-watcher/internal/service"
	"log"
	"time"
)

func main() {

	log.Println("Starting Metric Collector...")

	// 매트릭 프로세서 생성 (15초 간격으로 수집)
	processor := service.NewMetricProcessor(15 * time.Second)

	// 종료 채널 생성
	stopCh := make(chan struct{})

	// 메트릭 수집 시작
	processor.StartCollect(stopCh)

	log.Printf("Metric collector stopped // ")

}
