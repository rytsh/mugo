package funcs

import (
	"time"

	"github.com/rytsh/mugo/internal/render/generic"
)

func init() {
	generic.CallReg.AddFunction("time", generic.ReturnWithFn(Time{}))
}

type Time struct{}

func (Time) RFC3339() string {
	return "2006-01-02T15:04:05Z07:00"
}

func (Time) Format(format string, t time.Time) string {
	return t.Format(format)
}
