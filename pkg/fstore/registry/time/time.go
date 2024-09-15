package time

import (
	"time"
)

type Time struct{}

func (Time) Now() time.Time {
	return time.Now()
}

func (Time) RFC3339() string {
	return "2006-01-02T15:04:05Z07:00"
}

func (Time) Format(format string, t time.Time) string {
	return t.Format(format)
}

func (Time) UTC(t time.Time) time.Time {
	return t.UTC()
}

func (Time) AddDuration(t time.Time, d time.Duration) time.Time {
	return t.Add(d)
}

func (Time) Duration(d string) (time.Duration, error) {
	return time.ParseDuration(d)
}

func (Time) AddDate(t time.Time, years, months, days int) time.Time {
	return t.AddDate(years, months, days)
}
