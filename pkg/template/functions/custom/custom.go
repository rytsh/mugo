package custom

import (
	"github.com/spf13/cast"
)

func toUint64(v interface{}) uint64 {
	return cast.ToUint64(v)
}

func FuncMap() map[string]interface{} {
	fMap := map[string]interface{}{
		"uint64": toUint64,
	}

	return fMap
}
