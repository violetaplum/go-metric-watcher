package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

func GetEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func InitDB() (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	maxRetries := 5
	retryInterval := time.Second * 5

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Seoul",
		GetEnvOrDefault("DB_HOST", "localhost"),
		GetEnvOrDefault("DB_USER", "postgres"),
		GetEnvOrDefault("DB_PASSWORD", "password"),
		GetEnvOrDefault("DB_NAME", "metrics_db"),
		GetEnvOrDefault("DB_PORT", "5432"),
	)

	for i := 0; i < maxRetries; i++ {
		log.Printf("Attempting to connect to database (attempt %d/%d)", i+1, maxRetries)

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		})
		if err == nil {
			break
		}

		if i < maxRetries-1 {
			log.Printf("Failed to connect to database: %v. Retrying in %v...", err, retryInterval)
			time.Sleep(retryInterval)
		}
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database after %d attempts: %v", maxRetries, err)
	}

	// GORM의 자동 마이그레이션을 사용하지 않고 테이블이 이미 존재하는지만 확인
	var tableExists bool
	err = db.Raw("SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_name = 'alert_histories')").Scan(&tableExists).Error
	if err != nil {
		return nil, fmt.Errorf("failed to check if table exists: %v", err)
	}

	if !tableExists {
		log.Printf("Warning: alert_histories table does not exist. Please ensure init scripts are properly executed")
	}

	return db, nil
}
