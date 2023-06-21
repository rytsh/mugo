package templatex_test

import (
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

	output, err := tpl.ExecuteBuffer(
		templatex.WithData(map[string]interface{}{
			"a": 1,
			"b": 2,
		}),
		templatex.WithContent(`a + b = {{ add .a .b }}`+"\n"+`a - b = {{ sub .a .b }}`),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", output)
	// Output:
	// a + b = 3
	// a - b = -1
}
