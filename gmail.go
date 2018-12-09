package main

import (
	"net/smtp"
	"os"
)

// Gmail Получатель почты
type Gmail struct {
	Recipient string //Получатель
}

const (
	subject  string = "Subject: NotificationServer"
	smtpServ string = "smtp.gmail.com"
	smtpPort string = ":587"
)

// SendingMessGmail - Отправка сообщения через почту Gmail
// Использует аторизацию с помощью PlainAuth и отправку через
// SendMail(). Авторизация отправителя береться из .bashrc
func (g *Gmail) SendingMessGmail(mess string) {

	err := smtp.SendMail(smtpServ+smtpPort,
		smtp.PlainAuth("",
			os.Getenv("GmailUser"),
			os.Getenv("GmailPass"),
			smtpServ),
		os.Getenv("GmailUser"),
		[]string{g.Recipient},
		[]byte(subject+mess))

	if err != nil {
		Error.Println("smtp error: ", err)
		return
	}
}
