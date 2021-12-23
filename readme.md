# Tailpipe

<!--- mdtoc: toc begin -->

1. [Synopsis](#synopsis)
2. [Config](#config)
3. [Help](#help)<!--- mdtoc: toc end -->

## Synopsis

Tail a file and trigger an action if a new line occurs. Currently only mailing the new line somewhere is supported. Work in progress, there may be more to come...

## Config

Configuration file examples can be found in the `conf` folder.

This is how a mail config looks like:

```go mdox-exec="cat conf/mail.toml"
file_to_watch = "/tmp/tailpipe_test.log"
regex_filter = "(?i)(error|fail|fatal)"
action = "mail"

[mail]
host = "bear.local"
port = 1025
user = "test_user"
pass = "none"
encryption = "none"    # "none", "ssl", "tls"

addr_from = "tailpipe@vigilant.watch"
addr_to = ["recipient@mail.somewhere"]
subject = "Error on {{.Host}}"
template = """
Date:
{{.Date}}

Message:
{{.Text}}

{{if .TailError}}
There was a tail error:
{{.TailError}}
{{end}}
"""
```

## Help

```txt mdox-exec="r -h"

watch a text file for changes and do something with the latest line

Arguments:
  [<config-file>]    configuration file

Flags:
  -h, --help            Show context-sensitive help.
  -d, --debug           debug mode
  -V, --version-flag    display version
```
