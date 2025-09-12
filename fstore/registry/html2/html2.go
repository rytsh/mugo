package html2

import (
	"html"

	"github.com/rytsh/mugo/fstore"
)

func init() {
	fstore.AddStruct("html2", HTML2{})
}

type HTML2 struct{}

func (HTML2) EscapeString(v string) string {
	return html.EscapeString(v)
}

func (HTML2) UnescapeString(v string) string {
	return html.UnescapeString(v)
}
