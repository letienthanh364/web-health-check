package appCommon

import (
	"net/smtp"
	"os"
)

func SendEmail(to string, subject string, body string) error {
	password := os.Getenv("GO_SMTP_PWD")

	from := "letienthanh030604@gmail.com" // replace with your contact
	//password := "ebuf jcyp nqij ttyn" // replace with your contact password

	smtpHost := "smtp.gmail.com" // Gmail SMTP host
	smtpPort := "587"            // Gmail SMTP port for TLS

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(msg))
	if err != nil {
		return err
	}
	return nil
}
