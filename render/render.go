package render

import (
	"bytes"
	"errors"
	"log/slog"

	"github.com/rytsh/mugo/fstore"
	"github.com/rytsh/mugo/templatex"
)

var template = templatex.New(
	templatex.WithAddFuncMapWithOpts(func(o templatex.Option) map[string]any {
		return fstore.FuncMap(
			fstore.WithLog(slog.Default()),
			fstore.WithTrust(true),
			fstore.WithExecuteTemplate(o.T),
		)
	}),
)

var globalRender = Render{
	template: template,
}

type Render struct {
	template *templatex.Template
}

func (r *Render) Execute(content string) ([]byte, error) {
	return r.ExecuteWithData(content, nil)
}

func (r *Render) ExecuteWithData(content string, data any) ([]byte, error) {
	if r.template == nil {
		return nil, errors.New("template is nil")
	}

	var buf bytes.Buffer
	if err := r.template.Execute(
		templatex.WithIO(&buf),
		templatex.WithContent(content),
		templatex.WithData(data),
	); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func Execute(content string) ([]byte, error) {
	return globalRender.Execute(content)
}

func ExecuteWithData(content string, data any) ([]byte, error) {
	return globalRender.ExecuteWithData(content, data)
}
