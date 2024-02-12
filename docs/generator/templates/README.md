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

{{ template "BackToTop" }}

## ⌨️ Commands

{{ range . }}
{{- template "Command" . -}}
{{ end }}

---

## 📜 License

This software is licensed under the [MIT License](). See the [LICENSE](./LICENSE) file for details.

<!-- ----- -->
<!-- LINKS -->
<!-- ----- -->

[top]: #gh-license
