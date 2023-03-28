package sendEmailAlert

import (
	"fmt"

	"golang.org/x/exp/slog"
	"gopkg.in/gomail.v2"
	"io"
	"os"
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
func init() {
	logLevel := os.Getenv("LEVEL")
	//var level slog.Level
	var opt slog.HandlerOptions
	switch logLevel {
	case "Debug":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelDebug, // slog 默认日志级别是 info
		}
	case "Info":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelInfo, // slog 默认日志级别是 info
		}
	case "Warn":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelWarn, // slog 默认日志级别是 info
		}
	case "Err":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelError, // slog 默认日志级别是 info
		}
	default:
		slog.Warn("需要正确设置环境变量 Debug,Info,Warn or Err")
		slog.Info("默认使用Debug等级")
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelDebug, // slog 默认日志级别是 info
		}

	}
	file := "emailAlert.log"
	logf, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}
	defer logf.Close() //如果不关闭可能造成内存泄露
	logger := slog.New(opt.NewJSONHandler(io.MultiWriter(logf, os.Stdout)))
	slog.SetDefault(logger)
}

func Send(info *Info) {
	defer func() {
		if err := recover(); err != nil {
			slog.Warn("", slog.Any("发送邮件发生错误:%v\n", fmt.Sprint(err)))
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
		slog.Info("", slog.Any("", fmt.Sprintf("%+v", info)))
	}
	slog.Info("发送邮件", slog.Any("内容", info))
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
