.DEFAULT_GOAL := all

.PHONY: all
all: install lint build test

.PHONY: install
install:
	@go mod vendor

.PHONY: lint
lint:
	@golangci-lint run

.PHONY: format
format:
	@golangci-lint run --fix

.PHONY: build
build:
	@goreleaser build --clean --snapshot

.PHONY: test
test:
	@go test -race $(shell go list ./... | grep -v /vendor/)
