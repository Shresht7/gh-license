# `gh-license`

A GitHub CLI extension to view and generate license files.

All license information is obtained from the GitHub API (https://docs.github.com/en/rest/licenses).

> **NOTE**: The only purpose the CLI serves, is to save you the effort of opening up a browser to copy-&-paste the appropriate LICENSE text. It does NOT provide legal advice and you should still double-check the LICENSE yourself for any errors.

---

## 📦 Installation

### Requirements

- [GitHub CLI](https://cli.github.com/)

```sh
gh extension install Shresht7/gh-license
```

---

## 💻 Usage

Invoke the cli extension like so:

```sh
gh license <command>
```

For example, to generate a license file:

```sh
gh license create mit
```

![Usage Demonstration](docs/demo.gif)

<div align="right">

[⬆️ Back to top][top]

</div>

## ⌨️ Commands


### `create`

Create a license file

#### Alias

`new`, `add`, `init`, `set`

#### Usage

```sh
gh-license create [flags]
```

#### Flags

| Flag          | Description                                 | Default          |
| ------------- | ------------------------------------------- | ---------------- |
| `author`      | Author of the project                       | [AuthorName]     |
| `description` | Project description                         |                  |
| `output`      | Filepath                                    | LICENSE          |
| `project`     | Project name                                | [RepositoryName] |
| `web`         | Create license file using the web interface | false            |
| `year`        | Year                                        | [CurrentYear]    |

#### Examples

```sh
  gh-license create mit
  gh-license create mit --author Shresht7 --year 2023
  gh-license create --web
  gh-license create mit --web
```

<div align="right">

[⬆️ Back to top][top]

</div>


### `list`

Show a list of licenses



#### Usage

```sh
gh-license list [flags]
```

#### Flags

| Flag          | Description                     | Default |
| ------------- | ------------------------------- | ------- |
| `json`        | Output in JSON format           | false   |
| `pretty-json` | Output in pretty JSON format    | false   |
| `web`         | Open the license in the browser | false   |

#### Examples

```sh
  gh-license list
  gh-license list --json
  gh-license list --pretty-json
```

<div align="right">

[⬆️ Back to top][top]

</div>


### `repo`

View license of a repository

#### Alias

`r`

#### Usage

```sh
gh-license repo [flags]
```

#### Flags

| Flag          | Description                             | Default |
| ------------- | --------------------------------------- | ------- |
| `json`        | Print the license in JSON format        | false   |
| `pretty-json` | Print the license in pretty JSON format | false   |
| `web`         | Open the license in the browser         | false   |

#### Examples

```sh
  gh license repo
  gh license repo Shresht7/gh-license
  gh license repo gh-license
  gh license repo --json
  gh license repo Shresht7/gh-license --pretty-json
```

<div align="right">

[⬆️ Back to top][top]

</div>


### `view`

View details about a particular license

#### Alias

`show`, `get`

#### Usage

```sh
gh-license view [flags]
```

#### Flags

| Flag          | Description                                     | Default |
| ------------- | ----------------------------------------------- | ------- |
| `json`        | Print the license details in JSON format        | false   |
| `pretty-json` | Print the license details in pretty JSON format | false   |
| `web`         | Open the license in the browser                 | false   |

#### Examples

```sh
  gh-license view mit
  gh-license view mit --json
  gh-license view mit --pretty-json
```

<div align="right">

[⬆️ Back to top][top]

</div>



---

## 📜 License

This software is licensed under the [MIT License](). See the [LICENSE](./LICENSE) file for details.





[top]: #gh-license
