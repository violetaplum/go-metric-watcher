package notifier

import (
	"fmt"
	"net/smtp"
)

type EmailNotifier struct {
	from     string
	password string
	host     string
	port     string
}

func NewEmailNotifier(from, password, host, port string) *EmailNotifier {
	return &EmailNotifier{
		from:     from,
		password: password,
		host:     host,
		port:     port,
	}
}

func (e *EmailNotifier) Send(to []string, subject, body string) error {
	auth := smtp.PlainAuth("", e.from, e.password, e.host)
	msg := fmt.Sprintf("From: %s\r\n"+
		"To: %s\r\n"+
		"Subject: %s\r\n"+
		"\r\n"+
		"%s\r\n", e.from, to[0], subject, body)

	err := smtp.SendMail(e.host+":"+e.port, auth, e.from, to, []byte(msg))
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
