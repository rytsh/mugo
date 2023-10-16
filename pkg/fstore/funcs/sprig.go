package funcs

import (
	"github.com/Masterminds/sprig/v3"
	"github.com/rytsh/mugo/pkg/fstore/registry"
)

func init() {
	registry.AddGroup("sprig", sprig.GenericFuncMap)
}
