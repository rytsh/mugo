package sprig

import (
	"github.com/Masterminds/sprig/v3"

	"github.com/rytsh/mugo/fstore"
)

func init() {
	fstore.AddGroup("sprig", Sprig)
}

func Sprig() map[string]any {
	return sprig.GenericFuncMap()
}
