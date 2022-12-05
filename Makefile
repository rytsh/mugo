BINARY    := mugo
MAIN_FILE := cmd/$(BINARY)/main.go

PKG       := $(shell go list -m)
VERSION   := $(or $(IMAGE_TAG),$(shell git describe --tags --first-parent --match "v*" 2> /dev/null || echo v0.0.0))

LOCAL_BIN_DIR := $(PWD)/bin

## golangci configuration
GOLANGCI_CONFIG_URL   := https://raw.githubusercontent.com/worldline-go/guide/main/lint/.golangci.yml
GOLANGCI_LINT_VERSION := v1.50.0

.DEFAULT_GOAL := help

.PHONY: build golangci lint run test coverage clean help

run: ## Run the application
	go run $(MAIN_FILE)

build: ## Build the binary file
	goreleaser build --snapshot --rm-dist --single-target

.golangci.yml:
	@$(MAKE) golangci

golangci: ## Download .golangci.yml file
	@curl --insecure -o .golangci.yml -L'#' $(GOLANGCI_CONFIG_URL)

bin/golangci-lint-$(GOLANGCI_LINT_VERSION):
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(LOCAL_BIN_DIR) $(GOLANGCI_LINT_VERSION)
	@mv $(LOCAL_BIN_DIR)/golangci-lint $(LOCAL_BIN_DIR)/golangci-lint-$(GOLANGCI_LINT_VERSION)

lint: .golangci.yml bin/golangci-lint-$(GOLANGCI_LINT_VERSION) ## Lint Go files
	@$(LOCAL_BIN_DIR)/golangci-lint-$(GOLANGCI_LINT_VERSION) --version
	@GOPATH="$(shell dirname $(PWD))" $(LOCAL_BIN_DIR)/golangci-lint-$(GOLANGCI_LINT_VERSION) run ./...

test: ## Run unit tests
	@go test -v -race ./...

coverage: ## Run unit tests with coverage
	@go test -v -race -cover -coverpkg=./... -coverprofile=coverage.out -covermode=atomic ./...
	@go tool cover -func=coverage.out

html: ## Show html coverage result
	@go tool cover -html=./coverage.out

html-gen: ## Export html coverage result
	@go tool cover -html=./coverage.out -o ./coverage.html

html-wsl: html-gen ## Open html coverage result in wsl
	@explorer.exe `wslpath -w ./coverage.html` || true

clean: ## Remove previous build
	@rm -f $(BINARY)

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'