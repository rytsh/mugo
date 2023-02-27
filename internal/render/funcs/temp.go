package funcs

import (
	"github.com/rytsh/liz/utils/templatex"
	"github.com/rytsh/mugo/internal/render/generic"
)

func init() {
	generic.CallReg.AddFunction("execTemplate", new(ExecTemplate).init, "t")
}

type ExecTemplate struct {
	t *templatex.Template
}

func (e *ExecTemplate) init(t *templatex.Template) any {
	e.t = t

	return e.ExecTemplate
}

func (e *ExecTemplate) ExecTemplate(name string, v any) (string, error) {
	output, err := e.t.ExecuteBuffer(templatex.WithTemplate(name), templatex.WithData(v), templatex.WithParsed(true))
	return string(output), err
}
