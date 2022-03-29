package mail

import (
	"bytes"
	"crypto/rand"
	"embed"
	"fmt"
	"html/template"
	"path"
	"strings"
)

var (
	//go:embed templates
	tpl embed.FS
)

func (m *Mail) addStringToBody(str string) {
	m.Body.Write([]byte(str))
}

func (m *Mail) addTemplateToBody(str string) {
	data, err := tpl.ReadFile(path.Join("templates", str))
	if err != nil {
		fmt.Printf("mail template error: %s\n", err)
	}
	template.ParseFS(tpl, str)
	by := m.execTemplate(string(data))
	m.Body.Write(by.Bytes())
}

func (m *Mail) execTemplate(str string) bytes.Buffer {
	var buf bytes.Buffer
	tpl, err := template.New("template").Parse(str)
	if err != nil {
		fmt.Printf("[error] Can not parse mail template: %q\n", err)
		panic(err)
	}

	tpl.Execute(&buf, struct {
		Date           string
		Host           string
		TailError      string
		Text           string
		AddrFrom       string
		AddrTo         string
		Subject        string
		Boundary       string
		AttachFileName string
		AttachContent  string
		AttachMime     string
		UserAgent      string
	}{
		Date:           m.Payload.Date,
		Host:           m.Payload.Host,
		TailError:      m.Payload.TailErrorStr,
		Text:           m.Payload.Text,
		AddrFrom:       m.AddrFrom,
		AddrTo:         m.joinMailAddr(m.AddrTo),
		Subject:        m.Subject,
		Boundary:       m.Boundary,
		AttachFileName: m.AttachFileName,
		AttachContent:  m.AttachContent,
		AttachMime:     m.AttachMime,
		UserAgent:      "Tailpipe Mailer",
	})
	return buf
}

func (m *Mail) eraseBody() {
	m.Body.Reset()
}

func (m *Mail) print() {
	fmt.Printf("%s\n", string(m.Body.Bytes()))
}

func (m *Mail) joinMailAddr(arr []string) (r string) {
	var arr2 []string
	for _, el := range arr {
		arr2 = append(arr2, "<"+el+">")
	}
	r = strings.Join(arr2, ", ")
	return
}

func randStr(strSize int) string {
	dictionary := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var strBytes = make([]byte, strSize)
	_, _ = rand.Read(strBytes)
	for k, v := range strBytes {
		strBytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(strBytes)
}
