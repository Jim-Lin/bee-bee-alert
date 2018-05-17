package util

import (
	"net/smtp"
)

type Mail struct {
	Msg []byte
	To  []string
}

type MailTemplate struct {
	Subject string
	Msg     string
}

func (mt *MailTemplate) GetMail() *Mail {
	mail := GetConfig().MarketingMail

	return &Mail{
		To: []string{mail},
		Msg: []byte("To: " + mail + "\r\n" +
			"Subject: " + mt.Subject + "\r\n" +
			"\r\n" +
			mt.Msg),
	}
}

func (m *Mail) Notify() {
	var (
		hostname = GetConfig().SmtpHostname
		addr     = GetConfig().SmtpAddr
		from     = GetConfig().SmtpFrom
		password = GetConfig().SmtpPassword
	)

	auth := smtp.PlainAuth("", from, password, hostname)
	err := smtp.SendMail(addr, auth, from, m.To, m.Msg)
	checkMailError(err)
}

// slack...
// func (s *Slack) Notify() {
// }
