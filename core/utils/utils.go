package utils

import (
	"cloudDisk/core/define"
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"math/rand"
	"net/smtp"
	"time"

	"github.com/jordan-wright/email"
)

func Md5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func MailSendCode(mail string, code string) error {
	e := email.NewEmail()
	e.From = "Anon Tokyo <fireflycloud@yeah.net>"
	e.To = []string{mail}
	e.Subject = "验证码发送测试"
	e.HTML = []byte("<h1>你的验证码为：" + code + "</h1>")
	err := e.SendWithTLS("smtp.yeah.net:465",
		smtp.PlainAuth("", "fireflycloud@yeah.net", define.MailPass, "smtp.yeah.net"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.yeah.net"})
	if err != nil {
		return err
	}
	return nil
}
func RandomCode() string {
	s := "1234567890"
	code := ""
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < define.CodeLength; i++ {
		code += string(s[r.Intn(len(s))])
	}
	return code
}
