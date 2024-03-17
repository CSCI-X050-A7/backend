package email

import (
	"fmt"
	"net/smtp"

	"github.com/CSCI-X050-A7/backend/pkg/config"
)

func Send(to string, subject string, content string) error {
	conf := config.Conf
	username := conf.EmailUsername
	password := conf.EmailPassword
	smtpHost := conf.EmailSMTPHost
	smtpUrl := fmt.Sprintf("%s:%d", smtpHost, conf.EmailSMTPPort)
	auth := smtp.PlainAuth("", username, password, smtpHost)
	from := conf.EmailFrom
	toSlice := []string{to}
	message := []byte(fmt.Sprintf(
		"To: %s\r\nFrom: %s\r\nSubject: %s\r\n\r\n%s\r\n",
		to, from, subject, content,
	))
	err := smtp.SendMail(smtpUrl, auth, from, toSlice, message)
	return err
}
