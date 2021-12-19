package mail

type Mail struct {
	Host       string   `toml:"host"`
	Port       int      `toml:"port"`
	User       string   `toml:"user"`
	Pass       string   `toml:"pass"`
	Encryption string   `toml:"encryption"`
	AddrFrom   string   `toml:"addr_from"`
	AddrTo     []string `toml:"addr_to"`
	Subject    string   `toml:"subject"`
	Template   string   `toml:"template"`
}
