package templatex

import (
	"bytes"
	"fmt"
	"sort"
	textTemplate "text/template"

	"github.com/rytsh/mugo/pkg/templatex/store"
)

type Template struct {
	template       *textTemplate.Template
	templateParsed *textTemplate.Template
	funcs          map[string]interface{}
}

// New returns a new Template.
func New(opts ...store.Option) *Template {
	tpl := &Template{
		template: textTemplate.New("txt"),
	}

	tpl.setFunctions(opts...)

	return tpl
}

// ListFuncs returns the list of functions with name order.
func (t *Template) ListFuncs() []Info {
	funcs := FuncInfos(t.funcs)
	sort.Slice(funcs, func(i, j int) bool {
		return funcs[i].Name < funcs[j].Name
	})

	return funcs
}

// Reset the template.
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

func (t *Template) setFunctions(opts ...store.Option) {
	optsNew := make([]store.Option, 0, len(opts)+1)
	optsNew = append(optsNew, store.WithFnValue(t))
	optsNew = append(optsNew, opts...)

	t.funcs = store.New(optsNew...).Funcs()
	t.template.Funcs(t.funcs)
}

// AddFuncMap for extra textTemplate functions.
func (t *Template) AddFuncMap(funcMap textTemplate.FuncMap) {
	for k, v := range funcMap {
		t.funcs[k] = v
	}

	t.template.Funcs(funcMap)
}

// AddFunc for adding a func to the template.
func (t *Template) AddFunc(name string, fn interface{}) {
	t.funcs[name] = fn

	t.template.Funcs(textTemplate.FuncMap{name: fn})
}

// ParseGlob parses the template definitions in the files identified by the pattern.
func (t *Template) ParseGlob(pattern string) error {
	if pattern == "" {
		return nil
	}

	tpl, err := t.template.ParseGlob(pattern)
	if err != nil {
		return fmt.Errorf("Parse error: %w", err)
	}

	t.template = tpl

	return nil
}

// Parse content and set new template to parsed.
func (t *Template) Parse(content string) error {
	tpl, err := t.template.Clone()
	if err != nil {
		return fmt.Errorf("execute clone error: %w", err)
	}

	// Execute the template and write the output to the buffer
	tpl, err = tpl.Parse(content)
	if err != nil {
		return fmt.Errorf("Parse error: %w", err)
	}

	t.templateParsed = tpl

	return nil
}

// Execute the template and write the output to the buffer.
// Add WithIO to change the writer.
func (t *Template) Execute(opts ...Option) error {
	o := &options{
		writer: &bytes.Buffer{},
	}
	for _, opt := range opts {
		opt(o)
	}

	err := t.execute(o)
	if err != nil {
		return err
	}

	return nil
}

// ExecuteBuffer the template and return the output.
func (t *Template) ExecuteBuffer(opts ...Option) ([]byte, error) {
	output := &bytes.Buffer{}
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}

	o.writer = output

	err := t.execute(o)
	if err != nil {
		return nil, err
	}

	return output.Bytes(), nil
}

// ExecuteBuffer the template and return the output.
func (t *Template) ExecuteTemplate(templateName string, data any) ([]byte, error) {
	return t.ExecuteBuffer(WithTemplate(templateName), WithData(data), WithParsed(true))
}

func (t *Template) execute(o *options) error {
	tpl := t.template
	if o.parsed && t.templateParsed != nil {
		tpl = t.templateParsed
	}

	tpl, err := tpl.Clone()
	if err != nil {
		return fmt.Errorf("execute clone error: %w", err)
	}

	parsedTpl := tpl
	if !o.parsed {
		parsedTpl, err = tpl.Parse(o.content)
		if err != nil {
			return fmt.Errorf("execute parse error: %w", err)
		}
	}

	// Execute the template and write the output to the buffer
	if o.template != "" {
		err = parsedTpl.ExecuteTemplate(o.writer, o.template, o.data)
	} else {
		err = parsedTpl.Execute(o.writer, o.data)
	}

	if err != nil {
		return fmt.Errorf("Execute error: %w", err)
	}

	return nil
}
