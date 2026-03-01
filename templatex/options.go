package templatex

import "io"

type options struct {
	writer   io.Writer
	content  string
	template string
	data     any
	parsed   bool
	funcMap  map[string]any
}

// OptionExecute to execute the template.
type OptionExecute func(options *options)

// WithIO sets the writer to use.
// Useful for Execute function.
func WithIO(w io.Writer) OptionExecute {
	return func(options *options) {
		options.writer = w
	}
}

// WithContent sets the content to parse, if WithParsed used this option is ignored.
func WithContent(content string) OptionExecute {
	return func(options *options) {
		options.content = content
	}
}

// WithTemplate sets the specific template to execute.
func WithTemplate(template string) OptionExecute {
	return func(options *options) {
		options.template = template
	}
}

// WithData sets the data to use in Execute* functions.
// This is the values passed to the template.
func WithData(values any) OptionExecute {
	return func(options *options) {
		options.data = values
	}
}

// WithParsed sets the parsed template to use in Execute* functions.
func WithParsed(parsed bool) OptionExecute {
	return func(options *options) {
		options.parsed = parsed
	}
}

// WithExecFuncMap adds functions to the template for this execution only.
// These functions are applied to a cloned template and do not affect the base template.
func WithExecFuncMap(funcMap map[string]any) OptionExecute {
	return func(options *options) {
		if options.funcMap == nil {
			options.funcMap = make(map[string]any, len(funcMap))
		}

		for k, v := range funcMap {
			options.funcMap[k] = v
		}
	}
}

// WithExecFunc adds a single function to the template for this execution only.
// The function is applied to a cloned template and does not affect the base template.
func WithExecFunc(name string, fn any) OptionExecute {
	return func(options *options) {
		if options.funcMap == nil {
			options.funcMap = make(map[string]any, 1)
		}

		options.funcMap[name] = fn
	}
}
