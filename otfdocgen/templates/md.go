package templates

const (
	MD = `
| Name | Hex | Base64 |
|---|---|---|
{{range .Glyphs}}{{println "| " .Name " | " .HexString " | <img src=\"data:image/png;base64," .ImageBase64 "\" /> |"}}{{end}}
	`
)
