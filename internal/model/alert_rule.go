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
	ID             int64      `gorm:"column:id;primaryKey;autoIncrement"`
	Time           time.Time  `gorm:"column:time;primaryKey;not null"`
	AlertRuleID    int64      `gorm:"column:alert_rule_id;not null"`
	MetricName     string     `gorm:"column:metric_name;not null;type:text"`
	MetricValue    float64    `gorm:"column:metric_value;not null"`
	ThresholdValue float64    `gorm:"column:threshold_value;not null"`
	Status         string     `gorm:"column:status;not null;type:text"`
	Description    string     `gorm:"column:description;type:text"`
	ResolvedAt     *time.Time `gorm:"column:resolved_at"`
	TargetSystem   string     `gorm:"column:target_system;not null;type:text"`
	Severity       string     `gorm:"column:severity;not null;type:text"`
	CreatedAt      time.Time  `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
}

// TableName GORM 테이블 이름 지정
func (AlertHistory) TableName() string {
	return "alert_histories"
}
