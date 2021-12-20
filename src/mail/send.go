package mail

import (
	"bytes"
	"fmt"
	"net/smtp"
	"strconv"
	"strings"
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

	mimeHeaders := []string{
		"Content-Transfer-Encoding: quoted-printable",
		"Content-Type: text/plain; charset=UTF-8",
		"MIME-version: 1.0",
		"From: <" + conf.AddrFrom + ">",
		"To: " + joinMailAddr(conf.AddrTo),
		"Subject: " + conf.Subject,
		"\n\n",
	}
	body.Write([]byte(strings.Join(mimeHeaders, "\n")))

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

func joinMailAddr(arr []string) (r string) {
	var arr2 []string
	for _, el := range arr {
		arr2 = append(arr2, "<"+el+">")
	}
	r = strings.Join(arr2, ", ")
	return
}
