package appCommon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
	"os"
)

func SendEmail(to string, subject string, body string) error {
	password := os.Getenv("GO_SMTP_PWD")

	from := "letienthanh030604@gmail.com" // replace with your contact

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

func SendDiscordNotification(webhookURL, content string) error {
	// Create the payload
	payload := map[string]string{"content": content}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// Send the request
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check for a successful response
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to send notification, status code: %d", resp.StatusCode)
	}

	return nil
}
