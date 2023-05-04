# [Scrutinise](https://github.com/dbtedman/scrutinise) Contributing Guide

- [Core Maintainers](#core-maintainers)
- [Build](#build)
- [Lint and Format](#lint-and-format)
- [Test](#test)
- [Release](#release)
- [References](#references)

## Core Maintainers

- [Daniel Tedman](https://github.com/dbtedman)

## Build

```shell
make build
```

## Lint and Format

```shell
make lint
```

These rules can then be automatically applied:

```shell
make format
```

## Test

```shell
make test
```

## Release

> ⚠️ Creating new releases can only be completed by the [Core Maintainers](#core-maintainers).

Create a git tag and push to remote repository to trigger release workflow.

Once the release workflow has been completed, you will need to update the draft release that was created.

## References

> 💡 Resources referenced during the development of this project.

- [OSV Scanner (github.com)](https://github.com/google/osv-scanner) - Referenced for guidance on go release workflow.
