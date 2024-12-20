package notifier

import (
	"fmt"
	"github.com/violetaplum/go-metric-watcher/internal/model"
	"log"
	"os"
	"strings"
)

type AlertService struct {
	slackNotifier *SlackNotifier
	gmailNotifier *GmailNotifier
	thresholds    model.AlertThreshold
}

func NewAlertService(config *model.NotifierConfig) *AlertService {
	return &AlertService{
		slackNotifier: NewSlackNotifier(config.Slack.WebhookURL, config.Slack.Channel),
		gmailNotifier: NewGmailNotifier(config),
		thresholds:    config.Thresholds,
	}
}

func (a *AlertService) CheckMetricsAndAlert(metrics model.SystemMetric) error {
	var alerts []string

	if metrics.CPUUsage > a.thresholds.CPUUsage {
		alerts = append(alerts, fmt.Sprintf("CPU Usage Alert: %.2f%% (Threshold: %.2f%%)",
			metrics.CPUUsage, a.thresholds.CPUUsage))
	}

	if metrics.MemoryUsage > a.thresholds.MemoryUsage {
		alerts = append(alerts, fmt.Sprintf("Memory Usage Alert: %.2f%% (Threshold: %.2f%%)",
			metrics.MemoryUsage, a.thresholds.MemoryUsage))
	}

	if metrics.DiskUsage > a.thresholds.DiskUsage {
		alerts = append(alerts, fmt.Sprintf("Disk Usage Alert: %.2f%% (Threshold: %.2f%%)",
			metrics.DiskUsage, a.thresholds.DiskUsage))
	}

	if len(alerts) > 0 {
		message := strings.Join(alerts, "\n")
		if err := a.slackNotifier.Send(message); err != nil {
			log.Printf("Failed to send Slack alert: %v", err)
			log.Printf("Slack config: %v %v %v %v",
				os.Getenv("SLACK_WEBHOOK_URL"),
				os.Getenv("SLACK_CHANNEL"),
				os.Getenv("GMAIL_USER_NAME"),
				os.Getenv("GMAIL_PW"))
		}
		if err := a.gmailNotifier.Send(message); err != nil {
			log.Printf("Failed to send Gmail alert: %v", err)
		}
	}

	return nil
}
