package main

import (
	"net/smtp"
	"strings"
)

func main() {
	sendErrorMail("绑定", "1023")
}

func sendErrorMail(s string, eip string) {
	user := "tangshenzheng@163.com"
	nickname := "tangshenzheng"
	password := "forbidden920316"
	host := "smtp.163.com:25"
	to := []string{"754253197@qq.com"}
	content_type := "Content-Type: text/plain; charset=UTF-8"

	auth := smtp.PlainAuth("", user, password, "smtp.163.com")
	subject := s + "EIP: " + eip + "失败"
	body := s + "EIP: " + eip + "失败"
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	//msg := []byte(body)

	smtp.SendMail(host, auth, user, to, msg)
}
