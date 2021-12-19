package mail

import (
	"bytes"
	"fmt"
	"net/smtp"
	"strconv"
	"tailpipe/payload"
	"text/template"
)

// SendMail sends an email via smtp
func SendMail(p payload.Payload, conf Mail) {
	auth := smtp.PlainAuth("", conf.AddrFrom, conf.Pass, conf.Host)
	if conf.Encryption == "none" {
		auth = unencryptedAuth{
			smtp.PlainAuth(
				"", conf.AddrFrom, conf.Pass, conf.Host,
			),
		}
	}

	t, err := template.New("users").Parse(conf.Template)
	if err != nil {
		fmt.Printf("[error] Can not parse mail template: %q\n", err)
		panic(err)
	}
	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: "+conf.Subject+" \n%s\n\n", mimeHeaders)))
	errorStr := ""
	if p.TailError != nil {
		errorStr = fmt.Sprintf("%s", p.TailError)
	}
	t.Execute(&body, struct {
		Date      string
		Host      string
		TailError string
		Text      string
	}{
		Date:      p.Date,
		Host:      p.Host,
		TailError: errorStr,
		Text:      p.Text,
	})

	err = smtp.SendMail(
		conf.Host+":"+strconv.Itoa(conf.Port),
		auth, conf.AddrFrom, conf.AddrTo, body.Bytes(),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Email sent to via %s to %s\n", conf.Host, conf.AddrTo)
}
