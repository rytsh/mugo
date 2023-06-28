package funcs

import (
	"bytes"

	"github.com/rytsh/mugo/pkg/fstore/registry"
)

func init() {
	registry.AddFunction("execTemplate", new(ExecTemplate).init, "template")
}

type ExecTemplate struct {
	t registry.ExecuteTemplate
}

func (e *ExecTemplate) init(t registry.ExecuteTemplate) any {
	e.t = t

	return e.ExecTemplate
}

func (e *ExecTemplate) ExecTemplate(name string, v any) (string, error) {
	var buf bytes.Buffer
	err := e.t.ExecuteTemplate(&buf, name, v)
	return buf.String(), err
}
