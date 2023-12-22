package otpmailstorage

import (
	"fmt"
	"log"
	"net/smtp"
)

const (
	server   = "SMTP_SERVER_NAME"
	port     = "SMTP_PORT"
	user     = "SMTP_NAME"
	password = "SMTP_PASSWORD"
)

func (e *emailService) SendEmail(to, subject, body string) error {
	emailBody := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", to, subject, body)

	auth := smtp.PlainAuth("", e.User, e.Password, e.Server)

	err := smtp.SendMail(
		fmt.Sprintf("%s:%s", e.Server, e.Port),
		auth,
		e.User,
		[]string{to},
		[]byte(emailBody),
	)
	if err != nil {
		log.Panic("Error sending email:", err)
		return err
	}

	log.Print("Email sent successfully.")
	return nil
}
