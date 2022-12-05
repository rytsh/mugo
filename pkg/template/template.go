package template

import (
	"bytes"
	"fmt"
	"io"
	"sort"
	textTemplate "text/template"

	"github.com/rytsh/mugo/pkg/template/functions"
)

type Template struct {
	template *textTemplate.Template
	funcs    map[string]interface{}
}

func New() *Template {
	tpl := &Template{
		template: textTemplate.New("txt"),
	}

	tpl.setFunctions()

	return tpl
}

func (t *Template) ListFunctions() {
	funcs := FuncInfos(t.funcs)
	sort.Slice(funcs, func(i, j int) bool {
		return funcs[i].Name < funcs[j].Name
	})

	for _, v := range funcs {
		fmt.Println(v.Description)
	}
}

func (t *Template) Reset() {
	t.template = textTemplate.New("txt")
}

func (t *Template) SetDelims(left, right string) {
	if left == "" {
		left = "{{"
	}

	if right == "" {
		right = "}}"
	}

	t.template.Delims(left, right)
}

func (t *Template) setFunctions() {
	t.funcs = functions.Global.InitializeFuncs().Funcs()
	t.template.Funcs(t.funcs)
}

func (t *Template) Execute(v any, content string) (string, error) {
	var b bytes.Buffer
	// Execute the template and write the output to the buffer
	if err := textTemplate.Must(t.template.Parse(content)).Execute(&b, v); err != nil {
		return "", fmt.Errorf("Execute error: %w", err)
	}

	return b.String(), nil
}

func (t *Template) ExecuteContent(writer io.Writer, v any, content []byte) error {
	// Execute the template and write the output to the buffer
	if err := textTemplate.Must(t.template.Parse(string(content))).Execute(writer, v); err != nil {
		return fmt.Errorf("ExecuteContent error: %w", err)
	}

	return nil
}
