package render

import (
	"bytes"
	"errors"
	"log/slog"
	"sync"

	"github.com/rytsh/mugo/fstore"
	"github.com/rytsh/mugo/templatex"
)

var (
	globalRender *Render
	once         sync.Once
)

type Render struct {
	template *templatex.Template
}

func NewRender(opts ...fstore.OptionFunc) *Render {
	return &Render{
		template: templatex.New(
			templatex.WithAddFuncMapWithOpts(func(o templatex.Option) map[string]any {
				options := append([]fstore.OptionFunc{
					fstore.WithLog(slog.Default()),
					fstore.WithTrust(true),
					fstore.WithExecuteTemplate(o.T),
				}, opts...)

				return fstore.FuncMap(options...)
			}),
		),
	}
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
	once.Do(func() {
		globalRender = NewRender()
	})

	return globalRender.Execute(content)
}

func ExecuteWithData(content string, data any) ([]byte, error) {
	once.Do(func() {
		globalRender = NewRender()
	})

	return globalRender.ExecuteWithData(content, data)
}
