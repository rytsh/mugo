# templatex

`templatex` wraps Go's text and HTML template packages with reusable parsing and execution options.

```go
import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/rytsh/mugo/templatex"
)
```

## Usage

```go
tpl := templatex.New(
	templatex.WithAddFunc("upper", strings.ToUpper),
)

var output bytes.Buffer
if err := tpl.Execute(
	templatex.WithIO(&output),
	templatex.WithContent(`Hello {{ upper .Name }}`),
	templatex.WithData(map[string]any{"Name": "mugo"}),
); err != nil {
	log.Fatal(err)
}

fmt.Println(output.String())
// Output: Hello MUGO
```

Use `WithHTMLTemplate()` when automatic HTML escaping is required. Functions can also be added per execution with `WithExecFunc` or `WithExecFuncMap`.
