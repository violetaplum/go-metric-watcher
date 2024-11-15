package model

import watcherPb "github.com/violetaplum/go-metric-watcher/proto/gen/go/metrics/v1"

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

func (r *AlertRule) ToProto() *watcherPb.AlertRule {
	return &watcherPb.AlertRule{
		MetricType: r.MetricType,
		Threshold:  r.Threshold,
		Operator:   r.Operator,
		Duration:   r.Duration,
		Severity:   r.Severity,
		Channels:   r.Channels,
	}
}
