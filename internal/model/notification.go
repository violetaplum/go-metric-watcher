package model

import (
	"os"
)

type AlertThreshold struct {
	CPUUsage    float64 `json:"cpu_usage"`
	MemoryUsage float64 `json:"memory_usage"`
	DiskUsage   float64 `json:"disk_usage"`
}

type NotifierConfig struct {
	Slack struct {
		WebhookURL string `json:"webhook_url"`
		Channel    string `json:"channel"`
	} `json:"slack"`

	Gmail struct {
		Host     string   `json:"host"`
		Port     int      `json:"port"`
		Username string   `json:"username"`
		Password string   `json:"password"`
		To       []string `json:"to"`
	} `json:"gmail"`

	Thresholds AlertThreshold `json:"thresholds"`
}

func DefaultConfig() *NotifierConfig {
	config := &NotifierConfig{}

	// 기본값 설정
	config.Thresholds.CPUUsage = 1.0
	config.Thresholds.MemoryUsage = 1.0
	config.Thresholds.DiskUsage = 1.0

	config.Gmail.Host = "smtp.gmail.com"
	config.Gmail.Port = 587
	config.Gmail.Username = os.Getenv("GMAIL_USER_NAME")
	config.Gmail.Password = os.Getenv("GOOGLE_APP_PW")
	config.Gmail.To = []string{os.Getenv("GMAIL_TO")}

	config.Slack.WebhookURL = os.Getenv("SLACK_WEBHOOK_URL")
	config.Slack.Channel = os.Getenv("SLACK_CHANNEL")

	return config
}
