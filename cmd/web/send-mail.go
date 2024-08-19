package main

import (
	"log"

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
	go func() {
		err := sm.GetSendEmail().SendMail(m.From, m.To, m.Subject, string(m.Content))
		if err != nil {
			log.Printf("Error sending email: %v", err)
		}
	}()
}
