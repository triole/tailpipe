Date:
{{.Date}}

Message:
{{.Text}}

{{if .TailError}}
There was a tail error:
{{.TailError}}
{{end}}
