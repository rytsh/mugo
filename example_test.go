package mugo_test

import (
	"bytes"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/worldline-go/logz"

	"github.com/rytsh/mugo/fstore"
	"github.com/rytsh/mugo/templatex"
)

func Example() {
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
}
