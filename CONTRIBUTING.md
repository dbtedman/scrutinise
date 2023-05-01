# [Scrutinise](https://github.com/dbtedman/scrutinise) Contributing Guide

- [Build](#build)
- [Release](#release)

## Build

```shell
goreleaser build --clean --snapshot
```

## Release

Create a git tag and push to remote repository to trigger release workflow.

Once the release workflow has been completed, you will need to update the draft release that was created.
