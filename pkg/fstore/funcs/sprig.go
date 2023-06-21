//go:generate go run ./sprig_gen.go

package funcs

import (
	"github.com/Masterminds/sprig/v3"
	"github.com/rytsh/mugo/pkg/fstore/registry"
)

func init() {
	registry.DirectReg.Add("sprig", sprig.GenericFuncMap())
}
