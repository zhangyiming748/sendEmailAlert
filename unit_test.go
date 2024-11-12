package sendEmailAlert

import (
	"os"
	"testing"
	"time"
)

// go test -v -run  TestSend --timeout 5h59m -args ${{ secrets.PASSWORD }}
// args 从第五个参数开始
func TestSend(t *testing.T) {
	p := os.Getenv("PASSWD")
	info := new(Info)
	info.SetFrom("1914301892@qq.com") //${{ secrets.FROM }}
	tos := []string{
		"578779391@qq.com",
		"2352103020@qq.com",
		"zhangyiming748@qq.com",
		"zhangyiming7480@qq.com",
		"zhangyiming748@gmail.com",
		"zhangyiming748@protonmail.com",
		"zhangyiming748@outlook.com",
		"18904892728@163.com",
		"18904892728@189.cn",
		"zhangjializhangjiali@petalmail.com",
	}
	info.SetTo(tos)
	info.SetSubject("保活")
	info.SetText("发自github工作流")
	info.SetHost(QQ.SMTP)
	info.SetPort(QQ.SMTPProt)
	info.SetUsername("1914301892@qq.com") //${{ secrets.FROM }}
	info.SetPassword(p)                   //${{ secrets.PASSWORD }}
	info.AppendText(time.Now().String())
	status := info.Send()
	t.Log(status)
}
