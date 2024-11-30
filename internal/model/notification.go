package model

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
	config.Thresholds.CPUUsage = 80.0
	config.Thresholds.MemoryUsage = 85.0
	config.Thresholds.DiskUsage = 90.0

	config.Gmail.Host = "smtp.gmail.com"
	config.Gmail.Port = 587
	return config
}
