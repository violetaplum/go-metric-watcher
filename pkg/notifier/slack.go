package notifier

import (
	"bytes"
	"fmt"
	"github.com/goccy/go-json"
	"net/http"
)

type SlackNotifier struct {
	webhookURL string
	channel    string
}

func NewSlackNotifier(webhookURL, channel string) *SlackNotifier {
	return &SlackNotifier{
		webhookURL: webhookURL,
		channel:    channel,
	}
}

func (s *SlackNotifier) Send(message string) error {
	payload := map[string]interface{}{
		"channel": s.channel,
		"text":    message,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(s.webhookURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("slack notification failed with status: %d", resp.StatusCode)
	}
	return err
}
