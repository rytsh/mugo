package registry

import "fmt"

var (
	ErrInvalidType   = fmt.Errorf("invalid type")
	ErrTrustRequired = fmt.Errorf("trust required")
)
