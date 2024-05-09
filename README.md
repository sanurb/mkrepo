<div align = "center">

<h1><a href="https://github.com/sanurb/mkrepo">mkrepo</a></h1>

<a href="https://github.com/sanurb/mkrepo/blob/main/LICENSE">
<img alt="License" src="https://img.shields.io/github/license/sanurb/mkrepo?style=flat&color=eee&label="> </a>

<a href="https://github.com/sanurb/mkrepo/graphs/contributors">
<img alt="People" src="https://img.shields.io/github/contributors/sanurb/mkrepo?style=flat&color=ffaaf2&label=People"> </a>

<a href="https://github.com/sanurb/mkrepo/stargazers">
<img alt="Stars" src="https://img.shields.io/github/stars/sanurb/mkrepo?style=flat&color=98c379&label=Stars"></a>

<a href="https://github.com/sanurb/mkrepo/network/members">
<img alt="Forks" src="https://img.shields.io/github/forks/sanurb/mkrepo?style=flat&color=66a8e0&label=Forks"> </a>

<a href="https://github.com/sanurb/mkrepo/watchers">
<img alt="Watches" src="https://img.shields.io/github/watchers/sanurb/mkrepo?style=flat&color=f5d08b&label=Watches"> </a>

<a href="https://github.com/sanurb/mkrepo/pulse">
<img alt="Last Updated" src="https://img.shields.io/github/last-commit/sanurb/mkrepo?style=flat&color=e06c75&label="> </a>

<h3>Ready to go repos from the CLI 🏎️💨</h3>

</div>

**mkrepo** is a simple CLI tool that enhances the functionality of the GitHub CLI to streamline the creation of repositories. By handling setup tasks such as initializing the repository and setting up remotes, **mkrepo** allows developers to start coding immediately without worrying about initial setup chores.

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
## Table of Contents

- [✨ Features](#-features)
- [⚡ Setup](#-setup)
  - [⚙️ Requirements](#-requirements)
  - [💻 Installation](#-installation)
- [🚀 Usage](#-usage)
- [🧑‍💻 Behind The Code](#-behind-the-code)
  - [🌈 Inspiration](#-inspiration)
  - [💡 Challenges/Learnings](#-challengeslearnings)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## ✨ Features

- Automates the creation and setup of GitHub repositories.
- Configures remote origin for immediate push capability.
- Optional creation of public or private repositories.

## ⚡ Setup

### ⚙️ Requirements

- [Go](https://golang.org/dl/) 1.15 or higher
- [GitHub CLI](https://cli.github.com/)

### 💻 Installation

Install mkrepo using Go:

```bash
go install github.com/sanurb/mkrepo@latest
```

## 🚀 Usage

```bash
mkrepo <repoName> [options]
```


| Option         | Description                                                               | Default                         | Required |
|----------------|---------------------------------------------------------------------------|---------------------------------|----------|
| `<repoName>`   | The name of the new repository to create.                                 | None                            | Yes      |
| `[templateName]` | The name of the template repository to use.                              | `'bare-minimum'`                | No       |
| `[description]` | A brief description of the repository.                                   | `'Short Sweet Headline 🎇🎉'`   | No       |
| `--public`     | Include this flag to make the repository public. Omit for a private repo. | Private                         | No       |


## 🧑‍💻 Behind The Code

### 🌈 Inspiration

The inspiration for mkrepo came from the need to simplify repetitive tasks faced by developers when setting up new projects on GitHub.

### 💡 Challenges/Learnings

- Key learnings include mastering Go's CLI framework, Cobra

<hr>

<div align="center">

<strong>⭐ hit the star button if you found this useful ⭐</strong><br>

<a href="https://github.com/sanurb/mkrepo">Source</a>
| <a href="https://linkedin.com/in/sanurb" target="_blank">LinkedIn </a>
| <a href="https://sanurb.github.io/projects" target="_blank">Other Projects </a>

</div>