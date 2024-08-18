package models

import (
	"html/template"
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
	return &SMTPServer{}
}
