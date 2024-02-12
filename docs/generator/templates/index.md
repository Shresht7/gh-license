# `gh-license`

{{ .Description }}

## Description

{{ .LongDescription }}

## Usage

```
{{ .Examples }}
```

## Commands

{{ range .Commands }}
- [{{ .Name }}]({{ .Name }}.md) - {{ .Description }}{{ end }}
