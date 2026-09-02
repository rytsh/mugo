# Function Examples

These examples cover every registered Mugo function surface. The templates and expected outputs are executed by `go test` in [`fstore/documentation_test.go`](https://github.com/rytsh/mugo/blob/main/fstore/documentation_test.go) and `fstore/example_test.go`, so changes that break an example fail the test suite.

## Direct functions

Direct functions, including Sprig, do not require a namespace:

```txt
{{ nothing 1 }}{{ upper "mugo" }}
```

Output:

```txt
MUGO
```

## cast

```txt
{{ cast.ToInt "42" }}
```

Output: `42`

## codec

```txt
{{ $value := codec.JsonDecode (codec.StringToByte `{"name":"mugo"}`) -}}
{{ $value.name }}
```

Output: `mugo`

## crypto

```txt
{{ crypto.SHA256 "mugo" }}
```

Output: `db3e53a360a0f4ebebff1d40ef5b3305ce1928e6e31b91ac4184971c4e2388bc`

## exec

Command execution requires trust (`--trust` in the CLI or `fstore.WithTrust(true)` in Go):

```txt
{{ (exec.Exec "printf mugo").stdout | codec.ByteToString }}
```

Output: `mugo`

## execTemplate

```txt
{{ define "ochtend" }}Dag!{{ end }}{{ execTemplate "ochtend" nil | printf }}
```

Output: `Dag!`

## faker

Faker values are random, so this deterministic example verifies the UUID shape:

```txt
{{ len (faker.UUID.V4) }}
```

Output: `36`

## file

Writing requires trust. The automated test substitutes an isolated temporary path for `/tmp/mugo.txt`:

```txt
{{ file.Write "/tmp/mugo.txt" (codec.StringToByte "mugo") | nothing -}}
{{ file.Read "/tmp/mugo.txt" | codec.ByteToString }}
```

Output: `mugo`

## html2

```txt
{{ html2.EscapeString "<mugo>" }}
```

Output: `&lt;mugo&gt;`

## humanize

```txt
{{ humanize.Bytes 82854982 }}
```

Output: `83 MB`

## log

Log methods return their arguments, which allows a value to continue through a template pipeline:

```txt
{{ index (log.Info "rendered" "name" "mugo") 1 }}
```

Output: `mugo`

In Go, configure the destination with `fstore.WithLog(slog.Default())` or another `fstore.Adapter`.

## map

Slash-separated keys address nested maps:

```txt
{{ map.Set "app/name" "mugo" | nothing -}}
{{ map.Get "app/name" nil }}
```

Output: `mugo`

## math

The math functions use decimal arithmetic and accept numeric strings:

```txt
{{ math.Add "1.2" "2.3" }}
```

Output: `3.5`

## minify

```txt
{{ codec.ByteToString (minify "json" (codec.StringToByte `{ "name": "mugo" }`)) }}
```

Output: `{"name":"mugo"}`

## os

The automated test checks the temporary file created by the `file` example:

```txt
{{ os.FileExists "/tmp/mugo.txt" }}
```

Output: `true`

## random

Random values vary, so the stable example verifies the requested length:

```txt
{{ len (random.AlphaNum 8) }}
```

Output: `8`

Use `--random-seed` in the CLI when reproducible values are required.

## time

```txt
{{ time.Duration "2h" }}
```

Output: `2h0m0s`
