<div align="center">
  <img src="logo.png"/>
  <h1>PwGenie</h1>
  <h3>
    A simple password generator written in Golang.
  </h3>
  <p>
    <a href="CHANGELOG.md">
      <img
        alt="Keep a Changelog"
        src="https://img.shields.io/badge/changelog-Keep%20a%20Changelog-%23E05735"
      />
    </a>
    <a href="https://github.com/ntk148v/pwgenie/releases">
      <img
        alt="GitHub Release"
        src="https://img.shields.io/github/v/release/ntk148v/pwgenie"
      />
    </a>
    <a href="LICENSE">
      <img
        alt="GitHub License"
        src="https://img.shields.io/github/license/ntk148v/pwgenie"
      />
    </a>
    <a href="https://pkg.go.dev/github.com/ntk148v/pwgenie">
      <img
        alt="Go Release"
        src="https://pkg.go.dev/badge/github.com/ntk148v/pwgenie.svg"
      />
    </a>
    <a href="go.mod">
      <img
        alt="go.mod"
        src="https://img.shields.io/github/go-mod/go-version/ntk148v/pwgenie"
      />
    </a>
    <a
      href="https://github.com/ntk148v/pwgenie/actions?query=workflow%3Abuild+branch%3Amaster"
    >
      <img
        alt="Build Status"
        src="https://img.shields.io/github/actions/workflow/status/ntk148v/pwgenie/build.yml?branch=master"
      />
    </a>
    <br />
  </p>
  <br />
</div>

Table of content:

- [1. Introduction](#1-introduction)
- [2. Installation](#2-installation)
- [3. Usage](#3-usage)
- [4. Contributing](#4-contributing)

## 1. Introduction

PwGenie is a command-line application written in Golang to generate secure passwords.

It's highly inspired by [Motus](https://github.com/oleiade/motus) and [go-password](https://github.com/sethvargo/go-password/).

Features:

- Generate **secure human-friendly memorable passwords** using [EFF's wordlist](https://www.eff.org/deeplinks/2016/07/new-wordlists-random-passphrases).
- Generate **random passwords** with optional (uppercase, number, symbol inclusion), follow the algorithm described in [AgileBits 1Password](https://discussions.agilebits.com/discussion/23842/how-random-are-the-generated-passwords).
- Generate **PINs** with customizable length.
- Enable/disable **repeat**.
- **Clipboard** integration for easy password usage (Default).

## 2. Installation

## 3. Usage

```shell
$ pwgenie

pwgenie is a simple password generator.
<https://github.com/ntk148v/pwgenie>

Usage
-----

  pwgenie [OPTIONS] <SUBCOMMAND> [SUBCOMMAND-OPTIONS]

Options
-------

  -allow-repeat
                Allow repeat characters in the generated password

  -no-clipboard
                Disable automatic copying of generated password to clipboard

Subcommands
-----------

  human    Generate a human-friendly memorable password
  random   Generate a random password with specified complexity
  pin      Generate a random numeric PIN code

Run subcommand with '-h' for subcommand's options.

Example
-------

  $ pwgenie human
  trade clash striking underdog arbitrate

  $ pwgenie human -sep -
  preplan-mousiness-joining-eskimo-linguist

  $ pwgenie random
  bwuelvko

  $ pwgenie random -symb -num -upper
  _U*HkTzA
```

- Generate a human-friendly memorable password

```shell
$ pwgenie human -h

Generate a human-friendly memorable password

Usage of 'pwgenie human':
  -cap
        Enable capitalization of each word in the generated password
  -sep string
        The separator for words in the generated password (default " ")
  -words int
        The number of words in the generated password (default 5)

$ pwgenie human -sep - -words 10
stray-tableful-equity-stylishly-playmaker-pagan-upturned-gizzard-huntsman-defile

$ pwgenie human -cap
Daredevil Malt Recycler Prior Mutual
```

- Generate a random password

```shell
$ pwgenie random -h
Generate a random password with specified complexity

Usage of 'pwgenie random':
  -digit
        Enable the inclusion of numbers in the generated password
  -length int
        The number of characters in the generated password (default 8)
  -symbol
        Enable the inclusion of symbols in the generated password
  -upper
        Enable the inclusion of upper-case letters in the generated passwords

$ pwgenie random -digit -symbol -upper
iJdNmD0*

$ pwgenie random -digit -symbol -upper -length 20
LohapCbF_vzyuItDX91Z
```

- Generate a PIN

```shell
$ pwgenie pin -h
Generate a random numeric PIN code

Usage of 'pwgenie pin':
  -length int
        The number of digits in the generated PIN code (default 6)

$ pwgenie pin
491768
```

## 4. Contributing

We welcome contributions to the project. Feel free to submit issues, suggest new features, or create pull requests to help improve pwgenie.
