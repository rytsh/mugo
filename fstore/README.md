# fstore

`fstore` provides optional functions for Go templates.

```go
import (
	"bytes"
	"fmt"
	"log"
	"text/template"

	_ "github.com/rytsh/mugo/fstore/registry"

	"github.com/rytsh/mugo/fstore"
)
```

## Usage

The registry is opt-in. Import `fstore/registry` to register every function package, or blank-import selected packages under `fstore/registry/...`.

Sprig functions are added directly. Other functionality is generally exposed through structured entries such as `codec`, `file`, `math`, and `time`.

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

Functions that execute commands or write files require `fstore.WithTrust(true)`. Keep trust disabled when templates are not trusted.
