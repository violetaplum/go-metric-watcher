package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/violetaplum/go-metric-watcher/internal/model"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("this is api server...")

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")
	v1.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"time":   time.Now(),
		})
	})

	metrics := v1.Group("/metrics")

	metrics.GET("", func(c *gin.Context) {
		metrics := []model.SystemMetric{
			{
				Value:     123.45,
				Timestamp: time.Now(),
			},
			// 더 많은 메트릭 데이터..
		}
		c.JSON(http.StatusOK, metrics)
	})

	metrics.GET("/:name", func(c *gin.Context) {
		name := c.Param("name")
		metric := model.SystemMetric{
			Value:     123.45,
			Timestamp: time.Now(),
		}
		c.JSON(http.StatusOK, gin.H{
			"name":   name,
			"metric": metric,
		})
	})

	log.Println("API Server starting on :8080..")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
