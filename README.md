# gh-license
------------

A GitHub CLI extension to view and generate license files.

All license information is obtained from the GitHub API (https://docs.github.com/en/rest/licenses).

> **NOTE**: The only purpose the CLI serves is to save you the effort of opening up a browser to copy-&-paste the appropriate LICENSE text. It does NOT provide legal advice and you should still double-check the LICENSE yourself for any errors.

## Install

### Requirements

- GitHub CLI

```sh
gh extension install Shresht7/gh-license
```

## Usage

```sh
gh license <command>
```

### Help

```sh
gh license help
```

### View a list of all licenses

```sh
gh license list
```

### View details about a particular license

```sh
gh license view <licenseID>
```

### Create a license file

```sh
gh license create <licenseID>
```

#### Flags

| flag            | alias | description              | default      |
| --------------- | ----- | ------------------------ | ------------ |
| `--author`      | `-a`  | Specify the author name  |              |
| `--year`        | `-y`  | Specify the year         | Current year |
| `--project`     | `-p`  | Specify the project name |              |
| `--description` | `-d`  | Describe the project     |              |
| `--output`      | `-o`  | Output file              | `LICENSE`    |

#### Examples

##### Create a MIT license with author name

```sh
gh license create mit --author YourName
```

##### Create a GPL-3.0 license with project details in license.txt

```sh
gh license create GPL-3.0 --output license.txt --author YourName --year 2022 --project "Your Project" --description "An amazing project"
```
