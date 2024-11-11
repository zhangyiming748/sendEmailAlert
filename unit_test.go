package sendEmailAlert

import (
	"os"
	"testing"
)

// go test -v -run  TestSend --timeout 5h59m -args ${{ secrets.FROM }} ${{ secrets.TO }} 保活 发自github工作流 ${{ secrets.PASSWORD }}
// args 从第五个参数开始
func TestSend(t *testing.T) {
	args := os.Args
	for i, v := range args {
		t.Logf("第%d个参数:%s\n", i, v)
	}
	t.Log(args)
}
