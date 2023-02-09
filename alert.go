package sendEmailAlert

import (
	"github.com/zhangyiming748/log"
	"gopkg.in/gomail.v2"
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
	initInMain()
}

func initInMain() {
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	time.Local = cstZone
}

func Send(info *Info) {
	defer func() {
		if err := recover(); err != nil {
			log.Debug.Printf("发送邮件发生错误:%v\n", err)
		}
	}()
	m := gomail.NewMessage()
	m.SetHeader("From", info.Form)
	m.SetHeader("To", info.To...)
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", info.Subject)
	m.SetBody("text/html", info.Text)
	if info.Image != "" {
		m.Attach("/home/Alex/lolcat.jpg")
	}
	d := gomail.NewDialer(info.Host, info.Port, info.Username, info.Password)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	} else {
		log.Debug.Printf("%+v\n", info)
	}
}
func (i *Info) SetSubject(s string) {
	i.Subject = s
	return
}

//func (i Info) GetSubject() string {
//	return i.Subject
//}

func (i *Info) SetText(s string) {
	i.Text = s
	return
}

//	func (i Info) GetText() string {
//		return i.Text
//	}
func (i *Info) AddReceiver(s string) {
	i.To = append(i.To, s)
	return
}
func (i *Info) AddAllReceiver(s []string) {
	i.To = s
	return
}

//func (i Info) GetAllReceiver() []string {
//	return i.To
//}
