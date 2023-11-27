package funcs

import (
	"github.com/Masterminds/goutils"
	"github.com/rytsh/mugo/pkg/fstore/funcs/values"
	"github.com/rytsh/mugo/pkg/fstore/registry"
)

func init() {
	registry.AddGroup("random", registry.ReturnWithFn(Random{}))
}

type Random struct{}

func (Random) random(count int, start int, end int, letters bool, numbers bool, chars ...rune) (string, error) {
	return goutils.RandomSeed(count, start, end, letters, numbers, chars, values.RANDOM)
}

func (Random) Intn(min, max int) int {
	return values.RANDOM.Intn(max-min) + min
}
func (r Random) Alpha(n int) string {
	ret, _ := r.random(n, 0, 0, true, false)

	return ret
}

func (Random) AlphaNum(n int) string {
	ret, _ := Random{}.random(n, 0, 0, true, true)

	return ret
}
func (Random) Ascii(n int) string {
	ret, _ := Random{}.random(n, 32, 127, false, false)

	return ret
}

func (Random) Numeric(n int) string {
	ret, _ := Random{}.random(n, 0, 0, false, true)

	return ret
}

func (Random) Float(min float64, max float64) float64 {
	return values.RANDOM.Float64()*(max-min) + min
}
