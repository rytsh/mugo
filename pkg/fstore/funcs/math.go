package funcs

import (
	"math"

	"github.com/rytsh/mugo/pkg/fstore/registry"
)

func init() {
	registry.AddGroup("math", registry.ReturnWithFn(Math{}))
}

type Math struct{}

func (Math) RoundDecimal(precision int, value float64) float64 {
	return math.Round(value*math.Pow(10, float64(precision))) / math.Pow(10, float64(precision))
}
