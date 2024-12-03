package notifier

import (
	"context"
	"encoding/base64"
	"fmt"
	"golang.org/x/oauth2/google"
	gmail "google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
	"strings"
)

type GmailNotifier struct {
	service   *gmail.Service
	fromEmail string
}

func NewGmailNotifier(credentialJSON []byte, fromEmail string) (*GmailNotifier, error) {
	ctx := context.Background()

	credentials, err := google.CredentialsFromJSON(ctx, credentialJSON, gmail.GmailComposeScope)
	if err != nil {
		return nil, fmt.Errorf("failed to load credentials: %v", err)
	}

	service, err := gmail.NewService(ctx, option.WithCredentials(credentials))
	if err != nil {
		return nil, fmt.Errorf("failed to create Gmail service: %v", err)
	}

	return &GmailNotifier{service: service, fromEmail: fromEmail}, nil
}

func (g *GmailNotifier) Send(to []string, subject, body string) error {
	var message gmail.Message

	emailStr := fmt.Sprintf("From: %s\r\n"+
		"To: %s\r\n"+
		"Subject: %s\r\n"+
		"\r\n%s",
		g.fromEmail,
		strings.Join(to, ","),
		subject,
		body)

	message.Raw = base64.URLEncoding.EncodeToString([]byte(emailStr))

	_, err := g.service.Users.Messages.Send("me", &message).Do()
	if err != nil {
		return fmt.Errorf("[Send()] failed to send gmail: %v", err)
	}
	return nil
}
