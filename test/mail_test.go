package test

import (
	"cloudDisk/core/define"
	"crypto/tls"
	"net/smtp"
	"testing"

	"github.com/jordan-wright/email"
)

func TestMail(t *testing.T) {
	e := email.NewEmail()
	e.From = "Anon Tokyo <fireflycloud@yeah.net>"
	e.To = []string{"firefly@icloud.com"}
	e.Subject = "验证码发送测试"
	e.HTML = []byte("<h1>你的验证码为：123456</h1>")
	err := e.SendWithTLS("smtp.yeah.net:465",
		smtp.PlainAuth("", "firefly@yeah.net", define.MailPass, "smtp.yeah.net"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.yeah.net"})
	if err != nil {
		t.Fatal(err)
	}

}
