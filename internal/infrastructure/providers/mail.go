package providers

import (
	"errors"
	"senderEmails/internal"
	"senderEmails/internal/contracts"

	"gopkg.in/gomail.v2"
)

type MailProvider interface {
	SendMail(sendMailRequest contracts.SendMailRequest) error
}


type MailProviderImp struct {}


func (m *MailProviderImp) SendMail(sendMailRequest contracts.SendMailRequest) error {
	dialer := gomail.NewDialer(internal.MAIL_SMTP, internal.MAIL_PORT, internal.MAIL_FROM, internal.MAIL_PASSWORD)
	message := gomail.NewMessage()

	message.SetHeader("From", internal.MAIL_FROM)
	message.SetHeader("To", sendMailRequest.To...)
	message.SetHeader("Subject", sendMailRequest.Subject)
	message.SetBody("text/html", sendMailRequest.Message)


	err := dialer.DialAndSend(message); if err != nil {
		return errors.New("Erro ao enviar o email: " + err.Error())
	}

	return nil
}