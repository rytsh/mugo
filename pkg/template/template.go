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

// SetDelims sets the template delimiters to the specified strings
// and returns the template to allow chaining.
func (t *Template) SetDelims(left, right string) *Template {
	if left == "" {
		left = "{{"
	}

	if right == "" {
		right = "}}"
	}

	t.template.Delims(left, right)

	return t
}

func (t *Template) setFunctions() {
	t.funcs = functions.Global.InitializeFuncs().Funcs()
	t.template.Funcs(t.funcs)
}

func (t *Template) ParseGlob(pattern string) (*Template, error) {
	if pattern == "" {
		return t, nil
	}

	tpl, err := t.template.ParseGlob(pattern)
	if err != nil {
		return nil, fmt.Errorf("Parse error: %w", err)
	}

	t.template = tpl

	return t, nil
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
	// tpl, err := t.template.Clone()
	// if err != nil {
	// 	return fmt.Errorf("ExecuteContent clone error: %w", err)
	// }

	tpl := t.template
	// Execute the template and write the output to the buffer
	if err := textTemplate.Must(tpl.Parse(string(content))).Execute(writer, v); err != nil {
		return fmt.Errorf("ExecuteContent error: %w", err)
	}

	return nil
}
