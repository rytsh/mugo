package values

import (
	"math/rand"
	"time"
)

var RANDOM = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomSeed(seed int64) {
	RANDOM = rand.New(rand.NewSource(seed))
}
