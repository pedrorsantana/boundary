package smtp

import (
	"log"
	"net/smtp"
	"os"
)

func SendEmail(to string, subject string, body string) {
	//TODO: unmock
	from := os.Getenv("BOUNDARY_SMTP_FROM")
	pass := os.Getenv("BOUNDARY_SMTP_PASS")
	relay := os.Getenv("BOUNDARY_SMTP_RELAY")
	port := os.Getenv("BOUNDARY_SMTP_PORT")

	sendEmail(from, to, pass, subject, body, relay, port)
}

func sendEmail(from string, to string, pass string, subject string, body string, relay string, port string) {

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + " \n\n" +
		body
	err := smtp.SendMail(relay+":"+port,
		smtp.PlainAuth("", from, pass, relay),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
}
