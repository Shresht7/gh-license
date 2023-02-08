# `view`

View details about a particular license

## Aliases

`show`, `get`

## Description

View details about a particular license. For a list of available licenses, use the 'list' command.

## Usage

```sh
gh-license view [flags]
```

## Flags

| Flag | Description | Default |
|------|-------------|---------|
| `--json, -j` | Print the license details in JSON format | false |
| `--pretty-json, -p` | Print the license details in pretty JSON format | false |
| `--web, -w` | Open the license in the browser | false |

## Examples

```sh
  gh-license view mit
  gh-license view mit --json
  gh-license view mit --pretty-json
```

## See Also

* [`gh-license`](./gh-license.md)
* [`create`](./create.md)
* [`list`](./list.md)
* [`repo`](./repo.md)
