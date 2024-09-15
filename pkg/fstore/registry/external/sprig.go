package external

import (
	"github.com/Masterminds/sprig/v3"
)

func Sprig() map[string]interface{} {
	return sprig.GenericFuncMap()
}
