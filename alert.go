package sendEmailAlert

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"strings"
	"time"
)

type Info struct {
	Form     string   `json:"form"`
	To       []string `json:"to"`
	Subject  string   `json:"subject"`
	Text     string   `json:"text"`
	Image    string   `json:"image"`
	Host     string   `json:"host"`
	Port     int      `json:"port"`
	Username string   `json:"username"`
	Password string   `json:"password"`
}

func init() {
	initLocal()
}

func initLocal() {
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	time.Local = cstZone
}

func Send(info *Info) (status string) {
	defer func() {
		if err := recover(); err != nil {
			status = fmt.Sprintf("邮件发送失败%+v", err)
		} else {
			status = fmt.Sprintf("邮件发送成功%+v", info)
		}
	}()
	m := gomail.NewMessage()
	m.SetHeader("From", info.Form)
	m.SetHeader("To", info.To...)
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", info.Subject)
	m.SetBody("text/html", info.Text)
	if info.Image != "" {
		m.Attach(info.Image)
	}
	d := gomail.NewDialer(info.Host, info.Port, info.Username, info.Password)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	return status
}

func (i *Info) SetFrom(s string) {
	i.Form = s
}

func (i *Info) SetTo(s []string) {
	i.To = s
}

func (i *Info) SetSubject(s string) {
	i.Subject = s
}

func (i *Info) SetText(s string) {
	i.Text = s
}

/*
追加正文文本 已经添加了换行
*/

func (i *Info) AppendText(s string) {
	i.Text = strings.Join([]string{i.Text, s}, "<br>")
}

func (i *Info) SetImage(s string) {
	i.Image = s
}

func (i *Info) SetHost(s string) {
	i.Host = s
}

func (i *Info) SetPort(n int) {
	i.Port = n
}

func (i *Info) SetUsername(s string) {
	i.Username = s
}

func (i *Info) SetPassword(s string) {
	i.Password = s
}
