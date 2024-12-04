package template

import (
	"bytes"
	"io"
)

type ExecuteTemplate interface {
	ExecuteTemplate(wr io.Writer, name string, data any) error
}

type ExecTemplate struct {
	t ExecuteTemplate
}

func New(t ExecuteTemplate) *ExecTemplate {
	return &ExecTemplate{
		t: t,
	}
}

func (e *ExecTemplate) ExecTemplate(name string, v any) (string, error) {
	var buf bytes.Buffer
	err := e.t.ExecuteTemplate(&buf, name, v)
	return buf.String(), err
}
