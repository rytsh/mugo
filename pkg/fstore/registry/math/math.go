package math

import (
	"math"
)

type Math struct{}

func (Math) RoundDecimal(precision int, value float64) float64 {
	return math.Round(value*math.Pow(10, float64(precision))) / math.Pow(10, float64(precision))
}
