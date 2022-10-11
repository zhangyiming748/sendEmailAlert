package sendEmailAlert

import (
	"github.com/zhangyiming748/sendEmailAlert/log"
	"gopkg.in/gomail.v2"
	"time"
)

type Info struct {
	Form     string
	To       []string
	Subject  string
	Text     string
	Image    string
	Host     string
	Port     int
	Username string
	Password string
}

func init() {
	initInMain()
}

func initInMain() {
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	time.Local = cstZone
}

func Send(info *Info) {
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
