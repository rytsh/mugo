# Use in Go

Mugo's internal template runner and functions are available as a Go package.

## templatex

`templatex` is a package that provides a template runner with options.

```sh
go get github.com/rytsh/mugo
```

```sh
import "github.com/rytsh/mugo/templatex"
```

### Usage

Check details in go document: [https://pkg.go.dev/github.com/rytsh/mugo/templatex](https://pkg.go.dev/github.com/rytsh/mugo/templatex)

```go
tpl := templatex.New()

tpl.AddFunc("add", func(a, b int) int {
    return a + b
})

tpl.AddFunc("sub", func(a, b int) int {
    return a - b
})

var output bytes.Buffer
if err := tpl.Execute(
    templatex.WithIO(&output),
    templatex.WithData(map[string]interface{}{
        "a": 1,
        "b": 2,
    }),
    templatex.WithContent(`a + b = {{ add .a .b }}`+"\n"+`a - b = {{ sub .a .b }}`),
); err != nil {
    log.Fatal(err)
}

fmt.Printf("%s", output.String())
// Output:
// a + b = 3
// a - b = -1
```

## fstore

`fstore` is a package that provides bunch of functions with options.

```sh
go get github.com/rytsh/mugo
```

```sh
import "github.com/rytsh/mugo/fstore"
```

### Usage

Check details in go document: [https://pkg.go.dev/github.com/rytsh/mugo/fstore](https://pkg.go.dev/github.com/rytsh/mugo/fstore)

```go
tpl := template.New("test").Funcs(fstore.FuncMap())

output := &bytes.Buffer{}
tplParsed, err := tpl.Parse(`{{ $v := codec.JsonDecode (codec.StringToByte .) }}{{ $v.data.name }}`)
if err != nil {
    log.Fatal(err)
}

if err := tplParsed.Execute(output, `{"data": {"name": "Hatay"}}`); err != nil {
    log.Fatal(err)
}

fmt.Printf("%s", output)
// Output:
// Hatay
```

# Use fstore with templatex

`fstore` and `templatex` can be used together.
Use the tpl to execute templates.

```go
tpl := templatex.New(templatex.WithAddFuncsTpl(
    fstore.FuncMapTpl(
        fstore.WithLog(logz.AdapterKV{Log: log.Logger}),
        fstore.WithTrust(true),
        fstore.WithWorkDir("."),
    ),
))

var buf bytes.Buffer
err := tpl.Execute(
    templatex.WithContent("{{.Count}} items are made of {{.Material}}"),
    templatex.WithData(map[string]interface{}{
        "Count":    3,
        "Material": "wood",
    }),
    templatex.WithIO(&buf),
)
if err != nil {
    log.Fatal().Err(err).Msg("failed to execute template")
}

fmt.Println(buf.String())
// Output:
// 3 items are made of wood
```
