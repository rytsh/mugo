package templatex

import "io"

type options struct {
	writer   io.Writer
	content  string
	template string
	data     any
	parsed   bool
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
