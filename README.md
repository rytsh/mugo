![mugo](assets/mugo.svg)

[![License](https://img.shields.io/github/license/rytsh/mugo?color=red&style=flat-square)](https://raw.githubusercontent.com/rytsh/mugo/main/LICENSE)
[![Coverage](https://img.shields.io/sonar/coverage/rytsh_mugo?logo=sonarcloud&server=https%3A%2F%2Fsonarcloud.io&style=flat-square)](https://sonarcloud.io/summary/overall?id=rytsh_mugo)
[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/rytsh/mugo/Test?logo=github&style=flat-square&label=ci)](https://github.com/rytsh/mugo/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/rytsh/mugo?style=flat-square)](https://goreportcard.com/report/github.com/rytsh/mugo)
[![Go PKG](https://raw.githubusercontent.com/worldline-go/guide/main/badge/custom/reference.svg)](https://pkg.go.dev/github.com/rytsh/mugo)
<!-- [![Web](https://img.shields.io/badge/web-document-blueviolet?style=flat-square)](https://rytsh.github.io/mugo/) -->

Lightweight template executor. It is written in go and uses the [go template](https://golang.org/pkg/text/template/) package to render the templates.

Inspired by [hugo](https://gohugo.io/) to run small workflows.

## Usage

Input file could be anything which can include go template syntax.

```
Usage:
  mugo template.tpl [flags]

Examples:
mugo -d @data.yaml template.tpl
mugo -d '{"Name": "mugo"}' -o output.txt template.tpl

Flags:
  -d, --data stringArray   input data as json/yaml or file path with @ prefix
      --delims string      comma or space separated list of delimiters to alternate the default "{{ }}"
  -h, --help               help for mugo
  -l, --list               function List
  -o, --output string      output file, default is stdout
  -s, --silience           silience log
  -v, --version            version for mugo
```

### Development

<details><summary>Check details</summary>

Get binary with the goreleaser

```sh
make build
# goreleaser build --snapshot --rm-dist --single-target
```

</details>
