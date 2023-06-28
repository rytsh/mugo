package templatex

import (
	htmlTemplate "html/template"
	"io"
	"io/fs"
	textTemplate "text/template"
)

type templateInf interface {
	// AddParseTree(name string, tree *parse.Tree) (templateInf, error)
	Clone() (templateInf, error)
	DefinedTemplates() string
	Delims(left string, right string) templateInf
	Execute(wr io.Writer, data any) error
	ExecuteTemplate(wr io.Writer, name string, data any) error
	Funcs(funcMap map[string]any) templateInf
	Lookup(name string) templateInf
	Name() string
	New(name string) templateInf
	Option(opt ...string) templateInf
	Parse(text string) (templateInf, error)
	ParseFS(fsys fs.FS, patterns ...string) (templateInf, error)
	ParseFiles(filenames ...string) (templateInf, error)
	ParseGlob(pattern string) (templateInf, error)
	Templates() []templateInf
}

type templateX struct {
	templateText *textTemplate.Template
	templateHtml *htmlTemplate.Template

	isHtmlTemplate bool
}

func newTemplateX(name string, isHtml bool) templateX {
	if isHtml {
		return templateX{
			templateHtml:   htmlTemplate.New(name),
			isHtmlTemplate: true,
		}
	}

	return templateX{
		templateText: textTemplate.New(name),
	}
}

var _ templateInf = templateX{}

func (t templateX) Clone() (templateInf, error) {
	if t.isHtmlTemplate {
		tpl, err := t.templateHtml.Clone()
		if err != nil {
			return nil, err
		}

		t.templateHtml = tpl
	}

	tpl, err := t.templateText.Clone()
	if err != nil {
		return nil, err
	}

	t.templateText = tpl

	return t, nil
}

func (t templateX) DefinedTemplates() string {
	if t.isHtmlTemplate {
		return t.templateHtml.DefinedTemplates()
	}

	return t.templateText.DefinedTemplates()
}

func (t templateX) Delims(left string, right string) templateInf {
	if t.isHtmlTemplate {
		t.templateHtml.Delims(left, right)

		return t
	}

	t.templateText.Delims(left, right)

	return t
}

func (t templateX) Execute(wr io.Writer, data any) error {
	if t.isHtmlTemplate {
		return t.templateHtml.Execute(wr, data)
	}

	return t.templateText.Execute(wr, data)
}

func (t templateX) ExecuteTemplate(wr io.Writer, name string, data any) error {
	if t.isHtmlTemplate {
		return t.templateHtml.ExecuteTemplate(wr, name, data)
	}

	return t.templateText.ExecuteTemplate(wr, name, data)
}

func (t templateX) Funcs(funcMap map[string]any) templateInf {
	if t.isHtmlTemplate {
		t.templateHtml.Funcs(htmlTemplate.FuncMap(funcMap))

		return t
	}

	t.templateText.Funcs(textTemplate.FuncMap(funcMap))

	return t
}

func (t templateX) Lookup(name string) templateInf {
	if t.isHtmlTemplate {
		t.templateHtml = t.templateHtml.Lookup(name)

		return t
	}

	t.templateText = t.templateText.Lookup(name)

	return t
}

func (t templateX) Name() string {
	if t.isHtmlTemplate {
		return t.templateHtml.Name()
	}

	return t.templateText.Name()
}

func (t templateX) New(name string) templateInf {
	if t.isHtmlTemplate {
		t.templateHtml = t.templateHtml.New(name)
		return t
	}

	t.templateText = t.templateText.New(name)
	return t
}

func (t templateX) Option(opt ...string) templateInf {
	if t.isHtmlTemplate {
		t.templateHtml = t.templateHtml.Option(opt...)

		return t
	}

	t.templateText = t.templateText.Option(opt...)

	return t
}

func (t templateX) Parse(text string) (templateInf, error) {
	if t.isHtmlTemplate {
		tpl, err := t.templateHtml.Parse(text)
		if err != nil {
			return nil, err
		}

		t.templateHtml = tpl

		return t, nil
	}

	tpl, err := t.templateText.Parse(text)
	if err != nil {
		return nil, err
	}

	t.templateText = tpl

	return t, nil
}

func (t templateX) ParseFS(fsys fs.FS, patterns ...string) (templateInf, error) {
	if t.isHtmlTemplate {
		tpl, err := t.templateHtml.ParseFS(fsys, patterns...)
		if err != nil {
			return nil, err
		}

		t.templateHtml = tpl

		return t, nil
	}

	tpl, err := t.templateText.ParseFS(fsys, patterns...)
	if err != nil {
		return nil, err
	}

	t.templateText = tpl

	return t, nil
}

func (t templateX) ParseFiles(filenames ...string) (templateInf, error) {
	if t.isHtmlTemplate {
		tpl, err := t.templateHtml.ParseFiles(filenames...)
		if err != nil {
			return nil, err
		}

		t.templateHtml = tpl

		return t, nil
	}

	tpl, err := t.templateText.ParseFiles(filenames...)
	if err != nil {
		return nil, err
	}

	t.templateText = tpl

	return t, nil
}

func (t templateX) ParseGlob(pattern string) (templateInf, error) {
	if t.isHtmlTemplate {
		tpl, err := t.templateHtml.ParseGlob(pattern)
		if err != nil {
			return nil, err
		}

		t.templateHtml = tpl

		return t, nil
	}

	tpl, err := t.templateText.ParseGlob(pattern)
	if err != nil {
		return nil, err
	}

	t.templateText = tpl

	return t, nil
}

func (t templateX) Templates() []templateInf {
	if t.isHtmlTemplate {
		var templates []templateInf
		for _, tpl := range t.templateHtml.Templates() {
			templates = append(templates, templateX{
				templateHtml:   tpl,
				isHtmlTemplate: true,
			})
		}

		return templates
	}

	var templates []templateInf
	for _, tpl := range t.templateText.Templates() {
		templates = append(templates, templateX{
			templateText: tpl,
		})
	}

	return templates
}
