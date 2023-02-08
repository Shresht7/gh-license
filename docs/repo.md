# `repo`

View license of a repository

## Aliases

`r`

## Description

View license of a repository. If no repository is specified, the current repository is used.

## Usage

```sh
gh-license repo [flags]
```

## Flags

| Flag | Description | Default |
|------|-------------|---------|
| `--json, -j` | Print the license in JSON format | false |
| `--pretty-json, -p` | Print the license in pretty JSON format | false |
| `--web, -w` | Open the license in the browser | false |

## Examples

```sh
  gh license repo
  gh license repo Shresht7/gh-license
  gh license repo gh-license
  gh license repo --json
  gh license repo Shresht7/gh-license --pretty-json
```

## See Also

* [`gh-license`](./gh-license.md)
* [`create`](./create.md)
* [`list`](./list.md)
* [`view`](./view.md)
