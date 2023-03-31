# `{{ .Name }}`

{{ .Description }}

{{ if .Aliases -}}
## Aliases

{{ .Aliases }}
{{- end }}

{{ if .LongDescription -}}
## Description

{{ .LongDescription }}
{{- end }}

{{ if .Usage -}}
## Usage

```
{{ .Usage }}
```
{{- end }}

{{ if .Flags -}}
## Flags

{{ .Flags }}
{{- end }}

{{ if .Examples -}}
## Examples

```
{{ .Examples }}
```
{{ end }}
{{ if .SeeAlso -}}
## See Also

{{ range .SeeAlso -}}
- {{ . }}
{{ end }}
{{ end -}}
