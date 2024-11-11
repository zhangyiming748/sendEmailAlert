package sendEmailAlert

type Server struct {
	POP3     string
	POP3Port int
	SMTP     string
	SMTPProt int
}

var QQ = &Server{
	POP3:     "pop.qq.com",
	POP3Port: 995,
	SMTP:     "smtp.qq.com",
	SMTPProt: 465, //587
}

var NetEase = &Server{
	POP3:     "pop.163.com",
	POP3Port: 110,
	SMTP:     "smtp.163.com",
	SMTPProt: 25,
}

var Gmail = &Server{
	POP3:     "pop.gmail.com",
	POP3Port: 995,
	SMTP:     "smtp.gmail.com",
	SMTPProt: 587,
}

const (
	LF         = "\u000A"               // \n
	CR         = "\u000D"               // \r
	CRLf       = "\U000D000A"           // \r\n
	DOUBLECRLF = "\U000D000A\U000D000A" // \r\n\r\n
	NULL       = "\u0000"               // null
	NBSP       = "\u00A0"               // No-Break Space
	BR         = "<br>"
)
