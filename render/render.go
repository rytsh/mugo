package render

import (
	"bytes"
	"errors"
	"log/slog"

	"github.com/rytsh/mugo/fstore"
	"github.com/rytsh/mugo/templatex"
	"github.com/spf13/cast"
)

var Template = templatex.New(templatex.WithAddFuncsTpl(fstore.FuncMapTpl(
	fstore.WithLog(slog.Default()),
	fstore.WithTrust(true),
)))

var globalRender = Render{
	template: Template,
}

type Render struct {
	Data     map[string]any
	template *templatex.Template
}

func New() Render {
	return Render{
		template: Template,
	}
}

func (r *Render) Execute(content any) ([]byte, error) {
	return r.ExecuteWithData(content, r.Data)
}

func (r *Render) ExecuteWithData(content, data any) ([]byte, error) {
	if r.template == nil {
		return nil, errors.New("template is nil")
	}

	var buf bytes.Buffer
	if err := r.template.Execute(
		templatex.WithIO(&buf),
		templatex.WithContent(cast.ToString(content)),
		templatex.WithData(data),
	); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func Execute(content any) ([]byte, error) {
	return globalRender.Execute(content)
}

func ExecuteWithData(content, data any) ([]byte, error) {
	return globalRender.ExecuteWithData(content, data)
}
