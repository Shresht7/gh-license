# `create`

Create a license file

## Aliases

`new`, `add`, `init`, `set`

## Description

Create a license file for your project.

## Usage

```
gh-license create [flags]
```

## Flags

| Flag                | Type     | Description                                 | Default    |
| ------------------- | -------- | ------------------------------------------- | ---------- |
| `--author, -a`      | `string` | Author of the project                       | Shresht7   |
| `--description, -d` | `string` | Project description                         |            |
| `--output, -o`      | `string` | Filepath                                    | LICENSE    |
| `--project, -p`     | `string` | Project name                                | gh-license |
| `--web, -w`         | `bool`   | Create license file using the web interface | false      |
| `--year, -y`        | `string` | Year                                        | 2023       |

## Examples

```
  gh-license create mit
  gh-license create mit --author Shresht7 --year 2023
  gh-license create --web
  gh-license create mit --web

```

## See Also

* [gh-license](./gh-license.md)
* [list](./list.md)
* [repo](./repo.md)
* [view](./view.md)
