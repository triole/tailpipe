package mail

import (
	"bytes"
	"fmt"
	"net/smtp"
	"strconv"
	"strings"
	"tailpipe/src/payload"
	"text/template"
)

// SendMail sends an email via smtp
func SendMail(p payload.Payload, conf Mail) {
	if p.TailError != nil {
		p.TailErrorStr = fmt.Sprintf("%s", p.TailError)
	}

	auth := smtp.PlainAuth("", conf.AddrFrom, conf.Pass, conf.Host)
	if conf.Encryption == "none" {
		auth = unencryptedAuth{
			smtp.PlainAuth(
				"", conf.AddrFrom, conf.Pass, conf.Host,
			),
		}
	}

	mimeHeaders := []string{
		"Content-Transfer-Encoding: quoted-printable",
		"Content-Type: text/plain; charset=UTF-8",
		"MIME-version: 1.0",
		"From: <" + conf.AddrFrom + ">",
		"To: " + joinMailAddr(conf.AddrTo),
		"Subject: " + conf.Subject,
		"\n",
	}

	var body bytes.Buffer
	body.Write([]byte(strings.Join(mimeHeaders, "\n")))
	body.Write([]byte(conf.Template))
	body = execTemplate(body.String(), p)

	err := smtp.SendMail(
		conf.Host+":"+strconv.Itoa(conf.Port),
		auth, conf.AddrFrom, conf.AddrTo, body.Bytes(),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Email sent to via %s to %s\n", conf.Host, conf.AddrTo)
}

func joinMailAddr(arr []string) (r string) {
	var arr2 []string
	for _, el := range arr {
		arr2 = append(arr2, "<"+el+">")
	}
	r = strings.Join(arr2, ", ")
	return
}

func execTemplate(str string, p payload.Payload) bytes.Buffer {
	var buf bytes.Buffer
	tpl, err := template.New("template").Parse(str)
	if err != nil {
		fmt.Printf("[error] Can not parse mail template: %q\n", err)
		panic(err)
	}

	tpl.Execute(&buf, struct {
		Date      string
		Host      string
		TailError string
		Text      string
	}{
		Date:      p.Date,
		Host:      p.Host,
		TailError: p.TailErrorStr,
		Text:      p.Text,
	})
	return buf
}
