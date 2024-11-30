package notifier

import (
	"fmt"
	"github.com/violetaplum/go-metric-watcher/internal/model"
	"net/smtp"
)

type GmailNotifier struct {
	config model.NotifierConfig
}

func NewGmailNotifier(config model.NotifierConfig) *GmailNotifier {
	return &GmailNotifier{config: config}
}

func (g *GmailNotifier) Send(message string) error {
	auth := smtp.PlainAuth("", g.config.Gmail.Username,
		g.config.Gmail.Password, g.config.Gmail.Host)

	msg := fmt.Sprintf("Subject: System Alert\n\n%s", message)

	err := smtp.SendMail(fmt.Sprintf("%s:%d", g.config.Gmail.Host, g.config.Gmail.Port),
		auth,
		g.config.Gmail.Username,
		g.config.Gmail.To,
		[]byte(msg))
	return err
}
