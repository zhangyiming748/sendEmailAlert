# sendEmailAlert
通过电子邮件发送文件处理结果
# 用法
```shell
func TestSend(t *testing.T) {
	aim := []string{
		"xxx@example.com",
		"xxx@example.com",
	}
	var info = &Info{
		Form:     "",//发件人地址
		To:       aim,//收件人地址列表
		Subject:  "",//邮件主题
		Text:     "",//邮件正文
		Image:    "",//邮件附件
		Host:     "",//邮箱服务器
		Port:     ,//端口号
		Username: "",//发件人
		Password: "",//授权码
	}
	Send(info)
}
```