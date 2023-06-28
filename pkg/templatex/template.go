package templatex

import (
	"bytes"
	"fmt"
	"io"
	"sort"
	"sync"
)

var DefaultTemplateName = "_templatex"

type Template struct {
	template       templateInf
	templateParsed templateInf

	mutex sync.RWMutex
	funcs map[string]interface{}

	isHtmlTemplate bool
}

// New returns a new Template.
func New(opts ...OptionTemplate) *Template {
	tpl := &Template{}

	optsNew := make([]OptionTemplate, 0, len(opts)+1)
	optsNew = append(optsNew, WithFnValue(tpl))
	optsNew = append(optsNew, opts...)

	option := &optionsTemplate{}
	for _, opt := range optsNew {
		opt(option)
	}

	tpl.isHtmlTemplate = option.isHtmlTemplate
	tpl.template = newTemplateX(DefaultTemplateName, tpl.isHtmlTemplate)
	tpl.funcs = make(map[string]interface{}, len(option.addFuncs))

	tpl.AddFuncMap(
		option.addFuncs,
	)

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

// SetTypeHtml converts the template to html template.
//
// This function will reset the template when switching from text to html template.
func (t *Template) SetTypeHtml() {
	if !t.isHtmlTemplate {
		t.isHtmlTemplate = true
		t.Reset()
	}
}

// SetTypeText converts the template to text template.
//
// This function will reset the template when switching from html to text template.
func (t *Template) SetTypeText() {
	if t.isHtmlTemplate {
		t.isHtmlTemplate = false
		t.Reset()
	}
}

// Reset the template and add the functions back.
func (t *Template) Reset() {
	t.template = newTemplateX(DefaultTemplateName, t.isHtmlTemplate)
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

	for k, v := range funcMap {
		t.funcs[k] = v
	}

	t.template.Funcs(funcMap)
}

// AddFunc for adding a func to the template.
func (t *Template) AddFunc(name string, fn interface{}) {
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
