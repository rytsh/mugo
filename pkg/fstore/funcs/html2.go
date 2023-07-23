package funcs

import (
	"html"

	"github.com/rytsh/mugo/pkg/fstore/registry"
)

func init() {
	registry.AddGroup("html2", registry.ReturnWithFn(HTML2{}))
}

type HTML2 struct{}

func (HTML2) EscapeString(v string) string {
	return html.EscapeString(v)
}

func (HTML2) UnescapeString(v string) string {
	return html.UnescapeString(v)
}
