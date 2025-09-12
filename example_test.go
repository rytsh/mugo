package mugo_test

import (
	"bytes"
	"fmt"
	"log"

	_ "github.com/rytsh/mugo/fstore/registry"

	"github.com/rytsh/mugo/fstore"
	"github.com/rytsh/mugo/templatex"
)

func Example() {
	tpl := templatex.New(templatex.WithAddFuncMapWithOpts(func(o templatex.Option) map[string]any {
		return fstore.FuncMap(
			fstore.WithTrust(true),
			fstore.WithWorkDir("."),
			fstore.WithExecuteTemplate(o.T),
		)
	}))

	var buf bytes.Buffer
	err := tpl.Execute(
		templatex.WithContent("{{.Count}} items are made of {{.Material}}"),
		templatex.WithData(map[string]any{
			"Count":    3,
			"Material": "wood",
		}),
		templatex.WithIO(&buf),
	)
	if err != nil {
		log.Fatalf("failed to execute template: %v", err)
	}

	fmt.Println(buf.String())
	// Output:
	// 3 items are made of wood
}
