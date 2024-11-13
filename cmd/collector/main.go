package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("this is collector server...")
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// 1분마다 실행할 작업들 예시
			//collectCPUMetrics()
			//collectMemoryMetrics()
			//collectDiskMetrics()
			// ... 기타 메트릭 수집
			fmt.Println("10초가 지났습니다 메트릭수집을 시작합니다... 현재 시각은 :: ", time.Now())
		}
	}
}
