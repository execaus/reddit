package service

import (
	"gopkg.in/gomail.v2"
	"reddit/models"
)

const sender = "it-college-test-mail@yandex.ru"
const password = "KiD8qi9TbznzkCjQvCW0"
const smtpHost = "smtp.mail.ru"
const smtpPort = "465"

func sendEmailRegistration(data *models.InputSignUp) {
	message := "hello"
	sendEmail(message, data.Email)
}

func sendEmail(message string, receiver string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "alex@example.com")
	m.SetHeader("To", "bob@example.com", "cora@example.com")
	m.SetAddressHeader("Cc", receiver, "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", message)

	d := gomail.NewDialer("smtp.mail.ru", 465, sender, password)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
