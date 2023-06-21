package funcs

import (
	"html"

	"github.com/rytsh/mugo/pkg/fstore/registry"
)

func init() {
	registry.CallReg.AddFunction("html", registry.ReturnWithFn(Html{}))
}

type Html struct{}

func (Html) EscapeString(v string) string {
	return html.EscapeString(v)
}

func (Html) UnescapeString(v string) string {
	return html.UnescapeString(v)
}
