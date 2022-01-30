# gh-license
------------

A GitHub CLI extension to view and generate license files.

All license information is obtained from the GitHub API.

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
gh license view <license-id>
```

### Create a license file

```sh
gh license create <license-id>
```