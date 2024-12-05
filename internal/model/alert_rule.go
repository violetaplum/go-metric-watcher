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
	ID          uint       `json:"id"`
	AlertRuleID uint       `json:"alert_rule_id"`
	Status      string     `json:"status"`
	Description string     `json:"description"`
	Metric      float64    `json:"metric"`
	CreatedAt   time.Time  `json:"created_at"`
	ResolvedAt  *time.Time `json:"resolved_at,omitempty"`
}
