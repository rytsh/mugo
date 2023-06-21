package store

type options struct {
	disableFuncs []string
	addFuncs     map[string]interface{}
	fnValue      interface{}
}

type Option func(*options)

// WithDisableFuncs disables the functions.
func WithDisableFuncs(funcs ...string) Option {
	return func(o *options) {
		o.disableFuncs = append(o.disableFuncs, funcs...)
	}
}

// WithAddFuncs for adding multiple functions.
func WithAddFuncs(funcs map[string]interface{}) Option {
	return func(o *options) {
		if o.addFuncs == nil {
			o.addFuncs = make(map[string]interface{}, len(funcs))
		}

		for k, v := range funcs {
			o.addFuncs[k] = v
		}
	}
}

// WithAddFunc for adding a function.
func WithAddFunc(key string, f interface{}) Option {
	return func(o *options) {
		if o.addFuncs == nil {
			o.addFuncs = make(map[string]interface{}, 1)
		}

		o.addFuncs[key] = f
	}
}

// WithFnValue for passing a value to the function for WithAddFuncsTpl and WithAddFuncTpl.
// In templatex, default value is templatex.
func WithFnValue[T any](fn T) Option {
	return func(o *options) {
		o.fnValue = fn
	}
}

// WithAddFuncsTpl for adding multiple functions.
// The function is execute with the value passed to WithFnValue.
func WithAddFuncsTpl[T any](fn func(T) map[string]interface{}) Option {
	return func(o *options) {
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
func WithAddFuncTpl[T any](key string, f func(T) interface{}) Option {
	return func(o *options) {
		if o.addFuncs == nil {
			o.addFuncs = make(map[string]interface{}, 1)
		}

		o.addFuncs[key] = f(o.fnValue.(T))
	}
}
