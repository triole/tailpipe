package mail

import (
	"bytes"
	"tailpipe/src/payload"
)

type Mail struct {
	Host            string   `toml:"host"`
	Port            int      `toml:"port"`
	User            string   `toml:"user"`
	Pass            string   `toml:"pass"`
	Encryption      string   `toml:"encryption"`
	AddrFrom        string   `toml:"addr_from"`
	AddrTo          []string `toml:"addr_to"`
	Subject         string   `toml:"subject"`
	Template        string   `toml:"template"`
	AttachmentFiles []string
	Print           bool
	Body            bytes.Buffer
	Boundary        string
	Payload         payload.Payload
	AttachFileName  string
	AttachContent   string
	AttachMime      string
}
