package templatex

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"maps"
	"sort"
	"sync"
)

var DefaultTemplateName = "_templatex"

type Template struct {
	template       templateInf
	templateParsed templateInf

	mutex sync.RWMutex
	funcs map[string]any

	isHTMLTemplate bool
}

// New returns a new Template.
func New(opts ...OptionTemplate) *Template {
	tpl := &Template{}

	option := &optionsTemplate{
		opt: Option{T: tpl},
	}
	for _, opt := range opts {
		opt(option)
	}

	tpl.isHTMLTemplate = option.isHTMLTemplate
	tpl.template = newTemplateX(DefaultTemplateName, tpl.isHTMLTemplate)
	tpl.funcs = make(map[string]any, len(option.addFuncs))

	tpl.AddFuncMap(option.addFuncs)

	return tpl
}

// ListFuncs returns the list of functions with name order.
func (t *Template) ListFuncs() []Info {
	funcs := t.FuncInfos()
	sort.Slice(funcs, func(i, j int) bool {
		return funcs[i].Name < funcs[j].Name
	})

	return funcs
}

// SetTypeHTML converts the template to html template.
//
// This function will reset the template when switching from text to html template.
func (t *Template) SetTypeHTML() {
	if !t.isHTMLTemplate {
		t.isHTMLTemplate = true
		t.Reset()
	}
}

// SetTypeText converts the template to text template.
//
// This function will reset the template when switching from html to text template.
func (t *Template) SetTypeText() {
	if t.isHTMLTemplate {
		t.isHTMLTemplate = false
		t.Reset()
	}
}

// Reset the template and add the functions back.
func (t *Template) Reset() {
	t.template = newTemplateX(DefaultTemplateName, t.isHTMLTemplate)
	t.templateParsed = nil

	t.AddFuncMap(
		t.funcs,
	)
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

// AddFuncMap for extra textTemplate functions.
func (t *Template) AddFuncMap(funcMap map[string]any) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	maps.Copy(t.funcs, funcMap)

	t.template.Funcs(funcMap)
}

// AddFunc for adding a func to the template.
func (t *Template) AddFunc(name string, fn any) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.funcs[name] = fn

	t.template.Funcs(map[string]any{name: fn})
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

func (t *Template) ParseFS(fsys fs.FS, patterns ...string) error {
	if t.templateParsed == nil {
		tpl, err := t.template.Clone()
		if err != nil {
			return fmt.Errorf("execute clone error: %w", err)
		}

		t.templateParsed = tpl
	}

	tpl, err := t.templateParsed.ParseFS(fsys, patterns...)
	if err != nil {
		return fmt.Errorf("ParseFS error: %w", err)
	}

	t.templateParsed = tpl

	return nil
}

func (t *Template) Clone() (*Template, error) {
	tpl, err := t.template.Clone()
	if err != nil {
		return nil, fmt.Errorf("clone error: %w", err)
	}

	var tplParsed templateInf
	if t.templateParsed != nil {
		tplParsed, err = t.templateParsed.Clone()
		if err != nil {
			return nil, fmt.Errorf("clone parsed error: %w", err)
		}
	}

	return &Template{
		template:       tpl,
		templateParsed: tplParsed,
		funcs:          maps.Clone(t.funcs),
		isHTMLTemplate: t.isHTMLTemplate,
	}, nil
}

// Execute the template and write the output to the buffer.
// Add WithIO to change the writer.
//
// Example to execute a template with data and use parsed template:
//
//	t.Execute(WithTemplate(templateName), WithData(data), WithParsed(true))
//
// Example to execute and return the result:
//
//	var buf bytes.Buffer
//	t.Execute(WithIO(&buf), WithData(data))
func (t *Template) Execute(opts ...OptionExecute) error {
	o := &options{
		writer: &bytes.Buffer{},
	}
	for _, opt := range opts {
		opt(o)
	}

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

		t.templateParsed = parsedTpl
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

func (t *Template) ExecuteTemplate(wr io.Writer, name string, data any) error {
	return t.Execute(
		WithTemplate(name),
		WithData(data),
		WithIO(wr),
		WithParsed(true),
	)
}
