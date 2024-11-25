package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/violetaplum/go-metric-watcher/internal/model"
	"github.com/violetaplum/go-metric-watcher/pkg/monitoring"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("this is api server...")

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// API v1 그룹
	v1 := r.Group("/api/v1")

	// 헬스체크
	v1.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"time":   time.Now(),
		})
	})

	cpuMonitoring := monitoring.NewCPUMonitor()
	memoryMonitoring := monitoring.NewMemoryMonitor()
	diskMonitoring := monitoring.NewDiskMonitor("/")

	cpuMetrics, err := cpuMonitoring.Collect()
	if err != nil {
		log.Printf("Error on collecting cpu metrics: %v", err)
	}
	memoryMetrics, err := memoryMonitoring.Collect()
	if err != nil {
		log.Printf("Error on collecting memory metrics: %v", err)
	}
	diskMetrics, err := diskMonitoring.Collect()
	if err != nil {
		log.Printf("Error on collecting disk metrics: %v", err)
	}

	v1.GET("/metrics", func(c *gin.Context) {
		metrics := []model.SystemMetric{
			{
				Timestamp:   time.Now(),
				CPUUsage:    cpuMetrics.Usage,
				MemoryUsage: memoryMetrics.UsedPercent,
				MemoryTotal: memoryMetrics.Total,
				MemoryFree:  memoryMetrics.Free,
				DiskUsage:   diskMetrics.UsedPercent,
				DiskTotal:   diskMetrics.Total,
				DiskFree:    diskMetrics.Free,
			},
			// 더 많은 메트릭 데이터..
		}
		c.JSON(http.StatusOK, metrics)
	})

	log.Println("API Server starting on :8080..")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
