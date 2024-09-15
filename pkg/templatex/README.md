# templateX

templateX is a go template engine with some extra features.

```sh
import "github.com/rytsh/mugo/pkg/templatex"
```

## Usage

```go
tpl := templatex.New(store.WithAddFuncsTpl(
	fstore.FuncMapTpl(
		fstore.WithLog(logz.AdapterKV{Log: log.Logger}),
		fstore.WithTrust(true),
		fstore.WithWorkDir("."),
	),
))

tpl.Execute(
	templatex.WithIO(output),
	templatex.WithData(inputData),
	templatex.WithParsed(true),
);
```
