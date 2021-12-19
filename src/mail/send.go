package mail

import (
	"bytes"
	"fmt"
	"net/smtp"
	"strconv"
	"text/template"
)

// SendMail sends an email via smtp
func SendMail(msg string, conf Mail) {
	auth := smtp.PlainAuth("", conf.AddrFrom, conf.Pass, conf.Host)
	if conf.Encryption == "none" {
		auth = unencryptedAuth{
			smtp.PlainAuth(
				"", conf.AddrFrom, conf.Pass, conf.Host,
			),
		}
	}
	// auth = unencryptedAuth.Start(auth, auth.Server)

	t, err := template.New("users").Parse(conf.Template)
	if err != nil {
		fmt.Printf("[error] Can not parse mail template: %q\n", err)
		panic(err)
	}
	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: "+conf.Subject+" \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		Message string
	}{
		Message: msg,
	})

	err = smtp.SendMail(
		conf.Host+":"+strconv.Itoa(conf.Port),
		auth, conf.AddrFrom, conf.AddrTo, body.Bytes(),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent!")
}
