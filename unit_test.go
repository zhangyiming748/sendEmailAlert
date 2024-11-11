package sendEmailAlert

import (
	"os"
	"strings"
	"testing"
)

// go test -v -run  TestSend --timeout 5h59m -args ${{ secrets.FROM }} ${{ secrets.TO }} 保活 发自github工作流 ${{ secrets.PASSWORD }}
// args 从第五个参数开始
func TestSend(t *testing.T) {
	args := os.Args
	for i, v := range args {
		t.Logf("第%d个参数:%s\n", i, v)
	}
	info := new(Info)
	info.SetFrom(args[5])              //${{ secrets.FROM }}
	tos := strings.Split(args[6], ",") //${{ secrets.TO }}
	info.SetTo(tos)
	info.SetSubject(args[7]) //保活
	info.SetText(args[8])    //发自github工作流
	//info.SetImage(args[9])
	info.SetHost(QQ.SMTP)
	info.SetPort(QQ.SMTPProt)
	info.SetUsername(args[5]) //${{ secrets.FROM }}
	info.SetPassword(args[9]) //${{ secrets.PASSWORD }}
	t.Logf("%+v\n", info)
	status := info.Send()
	t.Log(status)
}
