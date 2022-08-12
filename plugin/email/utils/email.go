package utils

import (
	"Noteus/plugin/email/global"
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
	"strings"
)

func Email(To, subject string, body string) error {
	to := strings.Split(To, ",")
	return send(to, subject, body)
}

func ErrorToEmail(subject string, body string) error {
	to := strings.Split(global.GlobalConfig.To, ",")
	if to[len(to)-1] == "" {
		to = to[:len(to)-1]
	}
	return send(to, subject, body)
}

func EmailTest(subject string, body string) error {
	to := []string{global.GlobalConfig.From}
	return send(to, subject, body)
}

func send(to []string, subject string, body string) (err error) {
	form := global.GlobalConfig.From
	nickname := global.GlobalConfig.Nickname
	secret := global.GlobalConfig.Secret
	host := global.GlobalConfig.Host
	port := global.GlobalConfig.Port
	isSSL := global.GlobalConfig.IsSSL

	auth := smtp.PlainAuth("", form, secret, host)
	e := email.NewEmail()
	if nickname != "" {
		e.From = fmt.Sprintf("%s <%s>", nickname, form)
	} else {
		e.From = form
	}
	e.To = to
	e.Subject = subject
	e.HTML = []byte(body)
	hostAddr := fmt.Sprintf("%s:%d", host, port)
	if isSSL {
		err = e.SendWithTLS(hostAddr, auth, &tls.Config{ServerName: host})
	} else {
		err = e.Send(hostAddr, auth)
	}
	return
}
