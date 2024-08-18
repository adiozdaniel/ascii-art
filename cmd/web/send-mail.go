package main

import (
	"fmt"
	"time"

	"github.com/adiozdaniel/ascii-art/internals/models"
)

// listenForMail starts a goroutine to listen for email support requests and sends them via email
func listenForMail() {
	go func() {
		for msg := range sm.GetSupportChannel() {
			sendEmail(msg)
		}
	}()
}

// sendEmail sends an email using the provided data
func sendEmail(m *models.EmailData) {
	server := models.NewSMTPServer()
	server.Host = "smtp.gmail.com"
	server.Port = 587
	server.Username = "your-email@gmail.com"
	server.Password = "your-password"
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	err := server.SendMail(m.From, m.To, m.Subject, string(m.Content))
	if err != nil {
		// Handle the error, log it, or retry
		fmt.Printf("Failed to send email to %s: %v\n", m.To, err)
	} else {
		fmt.Printf("Email successfully sent to %s\n", m.To)
	}
}
