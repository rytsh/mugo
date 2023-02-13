package render

import (
	"time"
)

func rfc3339(t time.Time) string {
	return t.Truncate(time.Second).Format(time.RFC3339)
}
