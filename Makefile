BINARY    := mugo
MAIN_FILE := cmd/$(BINARY)/main.go

PKG       := $(shell go list -m)
VERSION   := $(or $(IMAGE_TAG),$(shell git describe --tags --first-parent --match "v*" 2> /dev/null || echo v0.0.0))

LOCAL_BIN_DIR := $(PWD)/bin

.DEFAULT_GOAL := help

.PHONY: run
run: ## Run the application
	go run $(MAIN_FILE)

.PHONY: docs-view
docs-view: ## Docs dev
	cd _documents && pnpm run docs:dev

.PHONY: test-data
test-data: ## Run the application
	go run $(MAIN_FILE) -d @testdata/data/input.yaml testdata/test.tpl

.PHONY: build
build: ## Build the binary file
	goreleaser build --snapshot --rm-dist --single-target

.PHONY: lint
lint: ## Lint Go files
	@golangci-lint --version
	@GOPATH="$(shell dirname $(PWD))" golangci-lint run ./...

.PHONY: lint-main
lint-main: ## Lint Go files with main branch diff
	@golangci-lint --version
	@GOPATH="$(shell dirname $(PWD))" golangci-lint run --new-from-rev main ./...

.PHONY: test
test: ## Run unit tests
	@go test -v -race -cover -coverpkg=./... -covermode=atomic ./...

.PHONY: help
help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
