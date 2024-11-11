package sendEmailAlert

import (
	"os"
	"testing"
)

// go test -v -run  TestSend --timeout 5h59m -args ${{ secrets.PASSWORD }}
// args 从第五个参数开始
func TestSend(t *testing.T) {
	args := os.Args
	for i, v := range args {
		t.Logf("第%d个参数:%s\n", i, v)
	}
	info := new(Info)
	info.SetFrom("1914301892@qq.com") //${{ secrets.FROM }}
	tos := []string{
		"578779391@qq.com",
	}
	info.SetTo(tos)
	info.SetSubject("保活")
	info.SetText("发自github工作流")
	info.SetHost(QQ.SMTP)
	info.SetPort(QQ.SMTPProt)
	info.SetUsername("1914301892@qq.com") //${{ secrets.FROM }}
	info.SetPassword(args[5])             //${{ secrets.PASSWORD }}
	t.Logf("%+v\n", info)
	status := info.Send()
	t.Log(status)
	t.Log(args)
}
