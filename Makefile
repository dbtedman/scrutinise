.DEFAULT_GOAL := all

.PHONY: all
all: install lint build test

.PHONY: install
install:
	@CGO_ENABLED=0 go mod vendor

.PHONY: lint
lint:
	@CGO_ENABLED=0 golangci-lint run

.PHONY: format
format:
	@CGO_ENABLED=0 golangci-lint run --fix

.PHONY: build
build:
	@CGO_ENABLED=0 goreleaser build --clean --snapshot

.PHONY: build_fast
build_fast:
	@CGO_ENABLED=0 go build -mod vendor -o ./dist/scrutinise ./

.PHONY: test
test:
	@CGO_ENABLED=1 go test -race $(shell go list ./... | grep -v /vendor/)
