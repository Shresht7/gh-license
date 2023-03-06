{{- define "Command" }}
### `{{ .Name }}`

{{ .Description }}

{{ if .Aliases -}}
#### Alias

{{ .Aliases }}
{{- end }}

{{ if .Usage -}}
#### Usage

```sh
{{ .Usage }}
```
{{- end }}

{{ if .Flags -}}
#### Flags

{{ .Flags }}
{{- end }}

{{ if .Examples -}}
#### Examples

```sh
{{ .Examples }}
```
{{- end }}

{{ template "BackToTop" }}

{{ end }}
