package main

import "github.com/adiozdaniel/ascii-art/internals/models"

// listenForMail starts a goroutine to listen for email support requests and sends them via email
func listenForMail() {
	go func() {
		for {
			msg := <-sm.GetSupportChannel()
			sendEmail(msg)
		}
	}()
}

// sendEmail sends an email using the provided data
func sendEmail(m *models.EmailData) {
	// Send email using m.To, m.Subject, and m.Body
}
