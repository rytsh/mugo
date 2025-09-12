package random

import (
	"math/rand"
	"time"

	"github.com/Masterminds/goutils"

	"github.com/rytsh/mugo/fstore"
)

var DefaultRandom = rand.New(rand.NewSource(time.Now().UnixNano()))

func SetDefaultRandomSeed(seed int64) {
	DefaultRandom.Seed(seed)
}

func init() {
	fstore.AddStruct("random", New(nil))
}

type Random struct {
	Random *rand.Rand
}

func New(r *rand.Rand) Random {
	if r == nil {
		r = DefaultRandom
	}

	return Random{
		Random: r,
	}
}

func (r Random) random(count int, start int, end int, letters bool, numbers bool, chars ...rune) (string, error) {
	return goutils.RandomSeed(count, start, end, letters, numbers, chars, r.Random)
}

func (r Random) Intn(min, max int) int {
	return r.Random.Intn(max-min) + min
}

func (r Random) Alpha(n int) string {
	ret, _ := r.random(n, 0, 0, true, false)

	return ret
}

func (r Random) AlphaNum(n int) string {
	ret, _ := r.random(n, 0, 0, true, true)

	return ret
}
func (r Random) Ascii(n int) string {
	ret, _ := r.random(n, 32, 127, false, false)

	return ret
}

func (r Random) Numeric(n int) string {
	ret, _ := r.random(n, 0, 0, false, true)

	return ret
}

func (r Random) Float(min, max float64) float64 {
	return r.Random.Float64()*(max-min) + min
}
