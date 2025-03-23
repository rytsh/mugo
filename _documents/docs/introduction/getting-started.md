# Getting Started

Mugo is a CLI tool to execute Go templates with a lot of functions to help you build your templates.

## Installation

### Binary

Download the binary from the [releases page](https://github.com/rytsh/mugo/releases/latest)

### Linux

```sh
curl -fSL https://github.com/rytsh/mugo/releases/latest/download/mugo_Linux_x86_64.tar.gz | tar -xz --overwrite -C ~/bin/ mugo
```

### Homebrew

Add the tap:

```sh
brew tap brew-tools/tap
```

Install the package:

```sh
brew install mugo
```

### Makefile

```makefile
.PHONY: check-tools
check-tools: ## Check if required tools are installed
	@echo "Checking if required tools are installed..."
	@command -v ~/bin/mugo > /dev/null || make tools
	@echo "All required tools are installed."

.PHONY: tools
tools: ## Install tools [mugo]
	@echo "Installing tools..."
	curl -fSL https://github.com/rytsh/mugo/releases/latest/download/mugo_Linux_x86_64.tar.gz | tar -xz --overwrite -C ~/bin/ mugo
```
