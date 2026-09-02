# Getting Started

Mugo is a CLI tool for rendering Go templates with structured or raw data and an extensible function library.

## Installation

### Binary

Download the archive for your operating system and architecture from the [releases page](https://github.com/rytsh/mugo/releases/latest).

### Linux

```sh
mkdir -p ~/bin
curl -fSL https://github.com/rytsh/mugo/releases/latest/download/mugo_Linux_x86_64.tar.gz | tar -xz --overwrite -C ~/bin/ mugo
```

Make sure `~/bin` is included in your `PATH`.

### macOS

```sh
mkdir -p ~/bin
curl -fSL https://github.com/rytsh/mugo/releases/latest/download/mugo_Darwin_arm64.tar.gz | tar -xz -C ~/bin/ mugo
```

Make sure `~/bin` is included in your `PATH`.

### Go

With Go 1.25 or newer:

```sh
go install github.com/rytsh/mugo/cmd/mugo@latest
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
	mkdir -p ~/bin
	curl -fSL https://github.com/rytsh/mugo/releases/latest/download/mugo_Linux_x86_64.tar.gz | tar -xz --overwrite -C ~/bin/ mugo
```
