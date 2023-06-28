package registry

import "io"

type ExecuteTemplate interface {
	ExecuteTemplate(wr io.Writer, name string, data any) error
}
