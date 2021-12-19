package main

import (
	"fmt"
	"os"

	"github.com/hpcloud/tail"

	tpm "tailpipe/mail"
)

func tailf(conf tConfig) {
	t, err := tail.TailFile(conf.FileToWatch, tail.Config{
		Follow:   true,
		Location: &tail.SeekInfo{0, os.SEEK_END},
	})
	if err != nil {
		fmt.Printf("%q\n", err)
	}
	for line := range t.Lines {
		if rxMatch(conf.RegexFilter, fmt.Sprintf("%s", line)) == true {
			tpm.SendMail(line.Text, conf.Mail)
		}
	}
}
