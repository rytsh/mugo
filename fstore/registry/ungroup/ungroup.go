package ungroup

import (
	"github.com/rytsh/mugo/fstore"
)

func init() {
	fstore.AddGroup("ungroup", Ungroup)
}

func Ungroup() map[string]any {
	return map[string]any{
		"nothing": Nothing,
	}
}

func Nothing(v ...any) string {
	return ""
}
