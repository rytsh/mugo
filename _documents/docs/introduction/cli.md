# CLI

Mugo accepts templates and data from command-line values, files, URLs, or standard input.

The positional argument or standard input is treated as the template by default. When `--template` is set, that value is the template and the positional argument or standard input is treated as data. `--data` supplies data explicitly; when both `--template` and `--data` are set, positional and standard input data are ignored.

Prefix a `--data` or `--template` value with `@` to load it from a local file. Use `--no-at` when a literal value must begin with `@`.

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
      --log-level string            log level (debug, info, warn, error), default is info (default "info")
      --no-at                       disable @ prefix for file path
      --no-retry                    disable retry on request
  -n, --no-stdin                    disable stdin input
  -o, --output string               output file, default is stdout
  -p, --parse stringArray           parse file pattern for define templates 'testdata/**/*.tpl'
      --perm-file string            create file permission, default is 0644
      --perm-folder string          create folder permission, default is 0755
      --random-seed int             seed for random function, default is 0 (random by time)
  -s, --silience                    silence log output
  -t, --template string             input template as raw or file path with @ prefix could be file with any extension
      --trust                       trust to execute dangerous functions
  -v, --version                     version for mugo
  -w, --work-dir string             work directory for run template
```

## Input data

Without `--data-raw`, inline data is decoded as YAML. JSON is valid YAML, so both formats can be passed directly. Files loaded with the `@` prefix support `.json`, `.yaml`, `.yml`, and `.toml` extensions.

Use `--data-raw` to pass data as a string. Add `--data-raw-byte` to expose file or input data as `[]byte` instead.

Multiple `--data` flags are merged in the order supplied.

## Templates

The positional template may be a local file, an HTTP(S) URL, or `-` for standard input. `--parse` may be repeated to parse additional template definitions before executing the main template.

Use `--html` to switch from `text/template` to `html/template`, and `--delims '<% %>'` to set custom left and right delimiters.

## Output and security

Rendered output is written to standard output unless `--output` is set. `--perm-folder` and `--perm-file` control permissions for newly created output paths.

Functions that execute commands or write files require `--trust`. Keep trust disabled for untrusted templates.

The silence flag is currently spelled `--silience` for CLI compatibility; `-s` is the recommended shorthand.
