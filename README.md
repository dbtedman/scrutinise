# [Scrutinise](https://github.com/dbtedman/scrutinise)

[![CI GitHub Pipeline](https://img.shields.io/github/actions/workflow/status/dbtedman/scrutinise/ci.yml?branch=main&style=for-the-badge&logo=github&label=ci)](https://github.com/dbtedman/scrutinise/actions/workflows/ci.yml?query=branch%3Amain)
[![SAST GitHub Pipeline](https://img.shields.io/github/actions/workflow/status/dbtedman/scrutinise/sast.yml?branch=main&style=for-the-badge&logo=github&label=sast)](https://github.com/dbtedman/scrutinise/actions/workflows/sast.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/dbtedman/scrutinise?style=for-the-badge)](https://goreportcard.com/report/github.com/dbtedman/scrutinise)

> ⚠️ WARNING! This project is in early development, and is not ready for production use.

Tool to scrutinise website development security.

-   [How to get started?](#how-to-get-started)
-   [How to contribute?](#how-to-contribute)
-   [Is this project secure?](#is-this-project-secure)
-   [License](#license)

## How to get started?

### Install (macOS)

```shell
brew tap dbtedman/tap
```

```shell
brew install scrutinise
```

### Install (release binary)

Download a pre-compiled binary from a [release (github.com)](https://github.com/dbtedman/scrutinise/releases) and install it into your [path (en.wikipedia.org)](https://en.wikipedia.org/wiki/PATH_(variable)_).

### Install (from source)

> Replacing `$VERSION` with the current latest version.

```shell
go install https://github.com/dbtedman/scrutinise@$VERSION
```

### Run

```shell
scrutinise
```

## How to contribute?

Read our [Contributing Guide](./CONTRIBUTING.md) to learn more about how to contribute to this project.

## Is this project secure?

Read our [Security Guide](./SECURITY.md) to learn how security is considered during the development and operation of this
tool.

## License

The [MIT License](./LICENSE.md) is used by this project.
