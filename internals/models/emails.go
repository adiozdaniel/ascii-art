package models

import (
	"bufio"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
	"strings"
	"time"
)

// EmailData represents data for sending email
type EmailData struct {
	Subject string
	Content template.HTML
	To      string
	From    string
}

// NewEmailData creates a new EmailData instance
func NewEmailData(subject, content, to, from string) *EmailData {
	return &EmailData{
		Subject: subject,
		Content: template.HTML(content),
		To:      to,
		From:    from,
	}
}

// SMTPServer holds configuration for the SMTP server
type SMTPServer struct {
	Host           string
	Port           int
	Username       string
	Password       string
	KeepAlive      bool
	ConnectTimeout time.Duration
	SendTimeout    time.Duration
}

// NewSMTPServer initializes and returns a new SMTPServer instance
func NewSMTPServer() *SMTPServer {
	return &SMTPServer{
		Host:           "smtp.gmail.com",
		Port:           587,
		Username:       os.Getenv("SMTP_USERNAME"),
		Password:       os.Getenv("SMTP_PASSWORD"),
		KeepAlive:      false,
		ConnectTimeout: 10 * time.Second,
		SendTimeout:    10 * time.Second,
	}
}

// SendMail sends an email using the configured SMTP server
func (s *SMTPServer) SendMail(from, to, subject, body string) error {
	auth := smtp.PlainAuth("", s.Username, s.Password, s.Host)

	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + body + "\r\n")

	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)
	return smtp.SendMail(addr, auth, from, []string{to}, msg)
}

// LoadEnv loads environment variables from a .env file
func (s *SMTPServer) LoadEnv() error {
	filename := instance.input.GetProjectRoot("./", ".env")
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		err := os.Setenv(key, value)
		if err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
