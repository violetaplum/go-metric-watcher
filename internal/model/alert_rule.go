package model

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
