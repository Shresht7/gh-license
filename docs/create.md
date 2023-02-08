# `create`

Create a license file

## Aliases

`new`, `add`, `init`, `set`

## Description

Create a license file for your project.

## Usage

```sh
gh-license create [flags]
```

## Flags

| Flag | Description | Default |
|------|-------------|---------|
| `--author, -a` | Author of the project | Shresht7 |
| `--description, -d` | Project description |  |
| `--output, -o` | Filepath | LICENSE |
| `--project, -p` | Project name | gh-license |
| `--web, -w` | Create license file using the web interface | false |
| `--year, -y` | Year | 2023 |

## Examples

```sh
  gh-license create mit
  gh-license create mit --author Shresht7 --year 2023
  gh-license create --web
  gh-license create mit --web
```

## See Also

* [`gh-license`](./gh-license.md)
* [`list`](./list.md)
* [`repo`](./repo.md)
* [`view`](./view.md)
