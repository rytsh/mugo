# CLI

```
Usage:
  mugo <template> [flags]

Examples:
mugo -d @data.yaml template.tpl
mugo -d '{"Name": "mugo"}' -o output.txt template.tpl
mugo -d '{"Name": "mugo"}' -o output.txt - < template.tpl
mugo -d '{"Name": "mugo"}' - <<< "{{.Name}}"

Flags:
  -d, --data stringArray            input data as json/yaml or file path with @ prefix could be '.yaml','.yml','.json','.toml' extension
  -r, --data-raw string             input data as raw or file path with @ prefix could be file with any extension
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
      --no-retry                    disable retry
  -o, --output string               output file, default is stdout
  -p, --parse stringArray           parse file pattern for define templates 'testdata/**/*.tpl'
      --perm-file string            create file permission, default is 0644
      --perm-folder string          create folder permission, default is 0755
  -s, --silience                    silience log
  -t, --trust                       trust to execute dangerous functions
  -v, --version                     version for mugo
  -w, --work-dir string             work directory for run template
```
