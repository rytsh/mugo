package templatex_test

import (
	"bytes"
	"fmt"
	"log"

	"github.com/rytsh/mugo/pkg/templatex"
)

func Example() {
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
}
