# fstore

fstore have bunch of go-template functions.

```sh
import "github.com/rytsh/mugo/fstore"
```

## Usage

__NOTE__ `sprig` functions added directly (direct group). Other functions added with struct.

Disable or just enable specific functions use options.

```go
tpl := template.New("test").Funcs(fstore.FuncMap())

output := &bytes.Buffer{}
tplParsed, err := tpl.Parse(`{{b64dec "TWVyaGFiYQ=="}}`)
if err != nil {
    log.Fatal(err)
}

if tplParsed.Execute(output, nil); err != nil {
    log.Fatal(err)
}

fmt.Printf("%s", output)
// Output:
// Merhaba
```
