![mugo](assets/mugo.svg)

[![License](https://img.shields.io/github/license/rytsh/mugo?color=red&style=flat-square)](https://raw.githubusercontent.com/rytsh/mugo/main/LICENSE)
[![Coverage](https://img.shields.io/sonar/coverage/rytsh_mugo?logo=sonarcloud&server=https%3A%2F%2Fsonarcloud.io&style=flat-square)](https://sonarcloud.io/summary/overall?id=rytsh_mugo)
[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/rytsh/mugo/test.yml?branch=main&logo=github&style=flat-square&label=ci)](https://github.com/rytsh/mugo/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/rytsh/mugo?style=flat-square)](https://goreportcard.com/report/github.com/rytsh/mugo)
[![Go PKG](https://raw.githubusercontent.com/rakunlabs/.github/main/assets/badges/gopkg.svg)](https://pkg.go.dev/github.com/rytsh/mugo)
[![Web](https://img.shields.io/badge/web-document-blueviolet?style=flat-square)](https://rytsh.github.io/mugo/)

Lightweight template executor. It is written in go and uses the [go template](https://golang.org/pkg/text/template/) package to render the templates.

Inspired by [hugo](https://gohugo.io/) to run small workflows.

## Usage

Input file could be anything which can include go template syntax.

```
Usage:
  mugo <template> [flags]

Examples:
mugo -d @data.yaml template.tpl
mugo -d '{"Name": "mugo"}' -o output.txt template.tpl
mugo -d '{"Name": "mugo"}' -o output.txt - < template.tpl
mugo -d '{"Name": "mugo"}' - <<< "{{.Name}}"
mugo -d '{"Name": "mugo"}' -t @template.tpl
mugo -t '{{.Name}}' data.yaml

Flags:
  -d, --data stringArray            input data as json/yaml or file path with @ prefix could be '.yaml','.yml','.json','.toml' extension
  -r, --data-raw                    set input data as raw
  -b, --data-raw-byte               raw data is byte
      --delims string               comma or space separated list of delimiters to alternate the default "{{ }}"
      --disable-func stringArray    disabled functions for run template
      --disable-group stringArray   disabled groups for run template
      --enable-func stringArray     specific functions for run template
      --enable-group stringArray    specific function groups for run template
  -h, --help                        help for mugo
      --html                        use html/template instead
  -k, --insecure                    skip verify ssl certificate
  -l, --list                        function list
      --log-level string            log level (debug, info, warn, error, fatal, panic), default is info (default "info")
      --no-at                       disable @ prefix for file path
      --no-retry                    disable retry on request
  -n, --no-stdin                    disable stdin input
  -o, --output string               output file, default is stdout
  -p, --parse stringArray           parse file pattern for define templates 'testdata/**/*.tpl'
      --perm-file string            create file permission, default is 0644
      --perm-folder string          create folder permission, default is 0755
      --random-seed int             seed for random function, default is 0 (random by time)
  -s, --silience                    silience log
  -t, --template string             input template as raw or file path with @ prefix could be file with any extension
      --trust                       trust to execute dangerous functions
  -v, --version                     version for mugo
  -w, --work-dir string             work directory for run template
```

## Example

Read all folder and generate info json seperately:

```sh
mugo -s -r -d "Moo-ve over" https://github.com/rytsh/mugo/raw/main/data/templates/cow.tpl
```

### Development

<details><summary>Build</summary>

Get binary with the goreleaser

```sh
make build
# goreleaser build --snapshot --rm-dist --single-target
```

</details>

<details><summary>Example</summary>

```sh
go run cmd/mugo/main.go -r -d "." -p 'testdata/tpl/*.tpl' - < testdata/readStart.tpl > output.json
go run cmd/mugo/main.go --trust -d '{"dir":"testdata","url":"http://localhost:5501"}'  -w "." - < testdata/readSeparate.tpl
go run cmd/mugo/main.go --trust -d '{"dir":"testdata","url":"http://localhost:5501", "output":"output"}'  -w "." data/templates/folderInfo.tpl
```

</details>
