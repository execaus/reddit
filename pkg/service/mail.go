package service

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"reddit/models"
)

const sender = "it-college-test-mail@yandex.ru"
const password = "nsgrbonartubistq"
const smtpHost = "smtp.yandex.ru"
const smtpPort = 465

func sendEmailRegistration(data *models.InputSignUp) {
	message := fmt.Sprintf("Вы успешно зарегистрировались в сервисе CustomReddit! %s:%s", data.Login, data.Password)
	sendEmail(message, data.Email)
}

func sendEmail(message string, receiver string) {
	m := gomail.NewMessage()
	m.SetHeader("From", sender)
	m.SetHeader("To", receiver)
	m.SetHeader("Subject", message)
	m.SetBody("text/html", message)

	d := gomail.NewDialer(smtpHost, smtpPort, sender, password)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
