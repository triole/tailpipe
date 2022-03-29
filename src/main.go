package main

import (
	"fmt"
	"path/filepath"
	tpm "tailpipe/src/mail"
	"tailpipe/src/payload"
)

func main() {
	parseArgs()

	confFile, _ := filepath.Abs(CLI.ConfigFile)
	conf := readConfig(confFile)

	if CLI.Mail != "" {
		tpm.SendMail(
			payload.NewTestPayload(CLI.Mail, CLI.Attachments),
			conf.Mail,
		)
	} else {
		fmt.Printf("Watch file %q\n", conf.FileToWatch)
		tailf(conf)
	}
}
