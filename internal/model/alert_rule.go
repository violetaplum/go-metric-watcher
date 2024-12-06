package model

import "time"

type AlertRule struct {
	ID          string
	MetricType  string
	Threshold   float64
	Operator    string
	Duration    int64
	Severity    string
	Channels    []string
	Description string
	Enabled     bool
}

type AlertHistory struct {
	ID             uint      `gorm:"primaryKey"`
	Time           time.Time `gorm:"index;not null"`
	AlertRuleID    int64     `gorm:"index;not null"`
	MetricName     string    `gorm:"type:varchar(100);not null"`
	MetricValue    float64   `gorm:"not null"`
	ThresholdValue float64   `gorm:"not null"`
	Status         string    `gorm:"type:varchar(20);not null"` // triggered, resolved
	Description    string    `gorm:"type:text"`
	ResolvedAt     *time.Time
	TargetSystem   string `gorm:"type:varchar(100);not null"`
	Severity       string `gorm:"type:varchar(20);not null"`
}
