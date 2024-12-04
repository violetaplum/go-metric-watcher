package notifier

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/violetaplum/go-metric-watcher/internal/model"
	"net/smtp"
	"os"
	"testing"
)

func Test_SMTP(t *testing.T) {
	from := os.Getenv("GMAIL_USER_NAME")
	password := os.Getenv("GOOGLE_APP_PW")

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte(fmt.Sprintf("To: %s\r\n"+
		"Subject: %s\r\n"+
		"\r\n"+
		"%s\r\n",
		"recipient@example.com", "테스트 이메일", "안녕하세요,\\n\\n이것은 Go로 보낸 테스트 이메일입니다."))

	auth := smtp.PlainAuth("", from, password, smtpHost)
	to := []string{os.Getenv("GMAIL_TO")}
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	assert.NoError(t, err)

	t.Log("Test Success!! ")
}

func Test_NewNotifierGmail(t *testing.T) {
	gmailNotifier := NewGmailNotifier(model.DefaultConfig())
	t.Log(gmailNotifier)
	err := gmailNotifier.Send("test test test ")
	assert.NoError(t, err)
}
