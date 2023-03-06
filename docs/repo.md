# `repo`

View license of a repository

## Aliases

`r`

## Description

View license of a repository. If no repository is specified, the current repository is used.

## Usage

```
gh-license repo [flags]
```

## Flags

| Flag                | Type   | Description                             | Default |
| ------------------- | ------ | --------------------------------------- | ------- |
| `--json, -j`        | `bool` | Print the license in JSON format        | false   |
| `--pretty-json, -p` | `bool` | Print the license in pretty JSON format | false   |
| `--web, -w`         | `bool` | Open the license in the browser         | false   |

## Examples

```
  gh license repo
  gh license repo Shresht7/gh-license
  gh license repo gh-license
  gh license repo --json
  gh license repo Shresht7/gh-license --pretty-json
```

## See Also

* [gh-license](./gh-license.md)
* [create](./create.md)
* [list](./list.md)
* [view](./view.md)
