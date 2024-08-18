package models

import "html/template"

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
