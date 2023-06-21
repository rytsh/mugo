package funcs

import (
	"math/big"
	"time"

	"github.com/dustin/go-humanize"

	"github.com/rytsh/mugo/pkg/fstore/registry"
)

func init() {
	registry.CallReg.AddFunction("humanize", registry.ReturnWithFn(Humanize{}))
}

// Humanize is a collection of humanize github.com/dustin/go-humanize functions.
type Humanize struct{}

func (Humanize) BigBytes(s *big.Int) string {
	return humanize.BigBytes(s)
}

func (Humanize) BigComma(b *big.Int) string {
	return humanize.BigComma(b)
}

func (Humanize) BigCommaf(v *big.Float) string {
	return humanize.BigCommaf(v)
}

func (Humanize) BigIBytes(s *big.Int) string {
	return humanize.BigIBytes(s)
}

func (Humanize) Bytes(v uint64) string {
	return humanize.Bytes(v)
}

func (Humanize) Comma(v int64) string {
	return humanize.Comma(v)
}

func (Humanize) Commaf(v float64) string {
	return humanize.Commaf(v)
}

func (Humanize) CommafWithDigits(f float64, decimals int) string {
	return humanize.CommafWithDigits(f, decimals)
}

func (Humanize) ComputeSI(input float64) (float64, string) {
	return humanize.ComputeSI(input)
}

func (Humanize) CustomRelTime(a time.Time, b time.Time, albl string, blbl string, magnitudes []humanize.RelTimeMagnitude) string {
	return humanize.CustomRelTime(a, b, albl, blbl, magnitudes)
}

func (Humanize) FormatFloat(format string, n float64) string {
	return humanize.FormatFloat(format, n)
}

func (Humanize) FormatInteger(format string, n int) string {
	return humanize.FormatInteger(format, n)
}

func (Humanize) Ftoa(num float64) string {
	return humanize.Ftoa(num)
}

func (Humanize) FtoaWithDigits(num float64, digits int) string {
	return humanize.FtoaWithDigits(num, digits)
}

func (Humanize) IBytes(v uint64) string {
	return humanize.IBytes(v)
}

func (Humanize) Ordinal(x int) string {
	return humanize.Ordinal(x)
}

func (Humanize) ParseBigBytes(s string) (*big.Int, error) {
	return humanize.ParseBigBytes(s)
}

func (Humanize) ParseBytes(s string) (uint64, error) {
	return humanize.ParseBytes(s)
}

func (Humanize) RelTime(a time.Time, b time.Time, albl string, blbl string) string {
	return humanize.RelTime(a, b, albl, blbl)
}

func (Humanize) SI(input float64, unit string) string {
	return humanize.SI(input, unit)
}

func (Humanize) Time(then time.Time) string {
	return humanize.Time(then)
}
