# `gh-license`

A GitHub CLI extension to view and generate license files.

All license information is obtained from the GitHub API (https://docs.github.com/en/rest/licenses).

> **NOTE**: The only purpose the CLI serves is to save you the effort of opening up a browser to copy-&-paste the appropriate LICENSE text. It does NOT provide legal advice and you should still double-check the LICENSE yourself for any errors.

<details>

<summary>Table of Contents</summary>

- [`gh-license`](#gh-license)
  - [📦 Installation](#-installation)
    - [Requirements](#requirements)
  - [💻 Usage](#-usage)
  - [⌨️ Commands](#️-commands)
    - [`help`](#help)
      - [Usage](#usage)
    - [`create`](#create)
      - [Alias](#alias)
      - [Usage](#usage-1)
      - [Flags](#flags)
      - [Examples](#examples)
        - [Create a MIT license with author name](#create-a-mit-license-with-author-name)
        - [Create a GPL-3.0 license with project details in license.txt](#create-a-gpl-30-license-with-project-details-in-licensetxt)
    - [`list`](#list)
      - [Usage](#usage-2)
    - [`view`](#view)
      - [Alias](#alias-1)
      - [Usage](#usage-3)
      - [Examples](#examples-1)
  - [📜 License](#-license)

</details>

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

To generate a license file:

```sh
gh license create MIT
```

<div align="right">

[⬆️ Back to top][top]

</div>

## ⌨️ Commands

### `help`

#### Usage

Use the `help` command to explore the cli.

```sh
gh license help
gh license help <command>
gh license <command> --help
```

<div align="right">

[⬆️ Back to top][top]

</div>

### `create`

#### Alias

`new`, `add`, `init`, `set`

#### Usage

Create a license file

```sh
gh license create <licenseID>
```

#### Flags

| Flag                | Description              | Default      |
| ------------------- | ------------------------ | ------------ |
| `--author, -a`      | Specify the author name  |              |
| `--year, -y`        | Specify the year         | Current year |
| `--project, -p`     | Specify the project name |              |
| `--description, -d` | Describe the project     |              |
| `--output, -o`      | Output file              | `LICENSE`    |

#### Examples

##### Create a MIT license with author name

```sh
gh license create mit --author YourName
```

##### Create a GPL-3.0 license with project details in license.txt

```sh
gh license create GPL-3.0 --output license.txt --author YourName --year 2022 --project "Your Project" --description "An amazing project"
```

<div align="right">

[⬆️ Back to top][top]

</div>

### `list`

#### Usage

View a list of all licenses

```sh
gh license list
```

<div align="right">

[⬆️ Back to top][top]

</div>

### `view`

#### Alias

`show`, `get`

#### Usage

View details about a particular license

```sh
gh license view <licenseID>
```

#### Examples

```sh
gh license view MIT
```

<div align="right">

[⬆️ Back to top][top]

</div>

---

## 📜 License

This software is licensed under the [MIT License](). See the [LICENSE](./LICENSE) file for details.

<!-- ----- -->
<!-- LINKS -->
<!-- ----- -->

[top]: #gh-license
