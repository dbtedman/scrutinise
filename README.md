# [Scrutinise](https://github.com/dbtedman/scrutinise)

[![CI GitHub Pipeline](https://img.shields.io/github/actions/workflow/status/dbtedman/scrutinise/ci.yml?branch=main&style=for-the-badge&logo=github&label=ci)](https://github.com/dbtedman/scrutinise/actions/workflows/ci.yml?query=branch%3Amain)
[![SAST GitHub Pipeline](https://img.shields.io/github/actions/workflow/status/dbtedman/scrutinise/sast.yml?branch=main&style=for-the-badge&logo=github&label=sast)](https://github.com/dbtedman/scrutinise/actions/workflows/sast.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/dbtedman/scrutinise?style=for-the-badge)](https://goreportcard.com/report/github.com/dbtedman/scrutinise)
[![GitHub all releases](https://img.shields.io/github/downloads/dbtedman/scrutinise/total?style=for-the-badge&logo=github)](https://github.com/dbtedman/scrutinise/releases)
[![Docker Pulls](https://img.shields.io/docker/pulls/dbtedman/scrutinise?sort=semver&style=for-the-badge&logo=docker)](https://hub.docker.com/r/dbtedman/scrutinise)

> ⚠️ WARNING! This project is in early development, and is not ready for production use.

Tool to scrutinise website development security.

| Supports  | Architecture    |
|-----------|-----------------|
| `linux`   | `amd64` `arm64` |
| `darwin`  | `amd64` `arm64` |
| `windows` | `amd64` `arm64` |

> 💡 `darwin` refers to the [macOS kernel (en.wikipedia.org)](https://en.wikipedia.org/wiki/Darwin_(operating_system)).

- [How to get started?](#how-to-get-started)
- [How to contribute?](#how-to-contribute)
- [Is this project secure?](#is-this-project-secure)
- [License](#license)

## How to get started?

- [Install (homebrew)](#install-homebrew)
- [Install (scoop)](#install-scoop)
- [Install (docker)](#install-docker)
- [Install (release binary)](#install-release-binary)
- [Install (from source)](#install-from-source)
- [Run](#run)

### Install (homebrew)

[github.com/dbtedman/homebrew-tap (github.com)](https://github.com/dbtedman/homebrew-tap/blob/main/Formula/scrutinise.rb)

```shell
brew tap dbtedman/tap && brew install scrutinise
```

### Install (scoop)

[github.com/dbtedman/scoop-bucket (github.com)](https://github.com/dbtedman/scoop-bucket/blob/main/bucket/scrutinise.json)

```shell
scoop bucket add dbtedman-scoop https://github.com/dbtedman/scoop-bucket
scoop install scrutinise
```

### Install (docker)

```shell
docker pull dbtedman/scrutinise
```

### Install (release binary)

Download a pre-compiled binary from a [release (github.com)](https://github.com/dbtedman/scrutinise/releases) and
install it into your [path (en.wikipedia.org)](https://en.wikipedia.org/wiki/PATH_(variable)_).

### Install (from source)

> Replacing `$VERSION` with the current latest version.

```shell
go install https://github.com/dbtedman/scrutinise@$VERSION
```

### Run

```shell
scrutinise
```

or via Docker

```shell
docker run -it dbtedman/scrutinise
```

## How to contribute?

Read our [Contributing Guide](./CONTRIBUTING.md) to learn more about how to contribute to this project.

## Is this project secure?

Read our [Security Guide](./SECURITY.md) to learn how security is considered during the development and operation of
this
tool.

## License

The [MIT License](./LICENSE.md) is used by this project.
