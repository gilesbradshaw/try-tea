# Gitea Command Line Tool for Go

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Release](https://raster.shields.io/badge/dynamic/json.svg?label=release&url=https://gitea.com/api/v1/repos/gitea/tea/releases&query=$[0].tag_name)](https://gitea.com/gitea/tea/releases)
[![Build Status](https://drone.gitea.com/api/badges/gitea/tea/status.svg)](https://drone.gitea.com/gitea/tea)
[![Join the chat at https://img.shields.io/discord/322538954119184384.svg](https://img.shields.io/discord/322538954119184384.svg)](https://discord.gg/Gitea)
[![Go Report Card](https://goreportcard.com/badge/code.gitea.io/tea)](https://goreportcard.com/report/code.gitea.io/tea)
[![GoDoc](https://godoc.org/code.gitea.io/tea?status.svg)](https://godoc.org/code.gitea.io/tea)

This project acts as a command line tool for operating one or multiple Gitea instances. It depends on [code.gitea.io/sdk](https://code.gitea.io/sdk) client SDK implementation written in Go to interact with
the Gitea API implementation.

## Installation

Currently no prebuilt binaries are provided.
To install, a Go installation is needed.

```sh
go get code.gitea.io/tea
go install code.gitea.io/tea
```

If the `tea` executable is not found, you might need to set up your `$GOPATH` and `$PATH` variables first:

```sh
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

If you have `brew` installed, you can install tea version via:

```sh
brew tap gitea/tap https://gitea.com/gitea/homebrew-gitea
brew install --devel tea
```

## Usage

First of all, you have to create a token on your `personal settings -> application` page of your gitea instance.
Use this token to login with `tea`:

```sh
tea login add --name=try --url=https://try.gitea.io --token=xxxxxx
```

Now you can use the `tea` commands:

```sh
tea issues
tea releases
```

To fetch issues from different repos, use the `--remote` flag (when inside a gitea repository directory) or `--login` & `--repo` flags.

## Compilation

To compile the sources yourself run the following:

```sh
go get code.gitea.io/tea
cd "${GOPATH}/src/code.gitea.io/tea"
go build
```

## Contributing

Fork -> Patch -> Push -> Pull Request

- `make test` run testsuite
- `make vendor` when adding new dependencies
- ... (for other development tasks, check the `Makefile`)

## Authors

* [Maintainers](https://github.com/orgs/go-gitea/people)
* [Contributors](https://github.com/go-gitea/tea/graphs/contributors)

## License

This project is under the MIT License. See the [LICENSE](LICENSE) file for the
full license text.
