package funcs

import (
	"github.com/rytsh/mugo/pkg/fstore/registry"
)

type ExecuteTemplate interface {
	ExecuteTemplate(templateName string, data any) ([]byte, error)
}

func init() {
	registry.CallReg.AddFunction("execTemplate", new(ExecTemplate).init, "template")
}

type ExecTemplate struct {
	t ExecuteTemplate
}

func (e *ExecTemplate) init(t ExecuteTemplate) any {
	e.t = t

	return e.ExecTemplate
}

func (e *ExecTemplate) ExecTemplate(name string, v any) (string, error) {
	output, err := e.t.ExecuteTemplate(name, v)
	return string(output), err
}
