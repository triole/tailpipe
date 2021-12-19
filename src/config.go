package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	tpm "tailpipe/mail"

	"github.com/pelletier/go-toml"
)

type tConfig struct {
	FileToWatch string `toml:"file_to_watch"`
	RegexFilter string `toml:"regex_filter"`
	Action      string `toml:"action"`
	Mail        tpm.Mail
}

// read the preset toml
func readConfig(filename string) (config tConfig) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading config: %s\n", err)
		os.Exit(1)
	}
	if err := toml.Unmarshal(content, &config); err != nil {
		fmt.Printf("Error decoding preset file: %q\n", err)
	}
	config.Mail.Encryption = strings.ToLower(config.Mail.Encryption)
	return
}
