package render

import (
	"time"
)

func Rfc3339(t time.Time) string {
	return t.Truncate(time.Second).Format(time.RFC3339)
}
