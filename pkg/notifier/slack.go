package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type SlackNotifier struct {
	webhookURL string
}

func NewSlackNotifier(webhookURL string) *SlackNotifier {
	return &SlackNotifier{webhookURL: webhookURL}
}

func (s *SlackNotifier) Send(message string) error {
	payload := map[string]string{
		"text": message,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal slack payload: %v", err)
	}

	resp, err := http.Post(s.webhookURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("failed to send slack notification: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("slack notification failed with status: %d", resp.StatusCode)
	}
	return nil
}
