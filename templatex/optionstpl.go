package templatex

type optionsTemplate struct {
	addFuncs map[string]interface{}
	fnValue  interface{}

	isHTMLTemplate bool
}

type OptionTemplate func(*optionsTemplate)

func WithHTMLTemplate() OptionTemplate {
	return func(o *optionsTemplate) {
		o.isHTMLTemplate = true
	}
}

// WithAddFuncMap for adding multiple functions.
func WithAddFuncMap(funcMap map[string]interface{}) OptionTemplate {
	return func(o *optionsTemplate) {
		if o.addFuncs == nil {
			o.addFuncs = make(map[string]interface{}, len(funcMap))
		}

		for k, v := range funcMap {
			o.addFuncs[k] = v
		}
	}
}

// WithAddFunc for adding a function.
func WithAddFunc(key string, f interface{}) OptionTemplate {
	return func(o *optionsTemplate) {
		if o.addFuncs == nil {
			o.addFuncs = make(map[string]interface{}, 1)
		}

		o.addFuncs[key] = f
	}
}

// WithFnValue for passing a value to the function for WithAddFuncsTpl and WithAddFuncTpl.
// In templatex, default value is templatex.
func WithFnValue[T any](fn T) OptionTemplate {
	return func(o *optionsTemplate) {
		o.fnValue = fn
	}
}

// WithAddFuncsTpl for adding multiple functions.
// The function is execute with the value passed to WithFnValue.
func WithAddFuncsTpl[T any](fn func(T) map[string]interface{}) OptionTemplate {
	return func(o *optionsTemplate) {
		funcs := fn(o.fnValue.(T))

		if o.addFuncs == nil {
			o.addFuncs = make(map[string]interface{}, len(funcs))
		}

		for k, v := range funcs {
			o.addFuncs[k] = v
		}
	}
}

// WithAddFuncTpl for adding a function.
// The function is execute with the value passed to WithFnValue.
func WithAddFuncTpl[T any](key string, f func(T) interface{}) OptionTemplate {
	return func(o *optionsTemplate) {
		if o.addFuncs == nil {
			o.addFuncs = make(map[string]interface{})
		}

		o.addFuncs[key] = f(o.fnValue.(T))
	}
}
