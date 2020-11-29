package email

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

type Email struct {
	*SMTPInfo
}
//smtpinfo协议服务器所包含的设置
type SMTPInfo struct {
	Host     string
	Port     int
	IsSSL    bool
	UserName string
	Password string
	From     string
}

//传入sm为email赋值
func NewEmail(info *SMTPInfo) *Email {
	return &Email{info}
}


//创建一个拨号器并发送邮件 传入拨号器的参数
func (e *Email) SendMail(to []string, subject, body string) error {
	m := gomail.NewMessage() //设置一个邮件 = 发送的内容
	m.SetHeader("From", e.From)
	m.SetHeader("TO", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	dialer := gomail.NewDialer(e.Host, e.Port, e.UserName, e.Password)//创建一个拨号装置
	dialer.TLSConfig=&tls.Config{InsecureSkipVerify:e.IsSSL}
	return dialer.DialAndSend(m) //向SMTP服务器发送一个连接 然后发送邮件m
}


//1.结构体
//2.new方法
//3.传入参数是结构体属性  然后做事情