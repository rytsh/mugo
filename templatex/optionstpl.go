package templatex

import "maps"

type Option struct {
	T *Template
}

type optionsTemplate struct {
	addFuncs map[string]any

	isHTMLTemplate bool
	opt            Option
}

type OptionTemplate func(*optionsTemplate)

func WithHTMLTemplate() OptionTemplate {
	return func(o *optionsTemplate) {
		o.isHTMLTemplate = true
	}
}

// WithAddFuncMap for adding multiple functions.
func WithAddFuncMap(funcMap map[string]any) OptionTemplate {
	return func(o *optionsTemplate) {
		if o.addFuncs == nil {
			o.addFuncs = make(map[string]any, len(funcMap))
		}

		for k, v := range funcMap {
			o.addFuncs[k] = v
		}
	}
}

// WithAddFunc for adding a function.
func WithAddFunc(key string, f any) OptionTemplate {
	return func(o *optionsTemplate) {
		if o.addFuncs == nil {
			o.addFuncs = make(map[string]any, 1)
		}

		o.addFuncs[key] = f
	}
}

func WithAddFuncWithOpts(fn func(o Option) (string, any)) OptionTemplate {
	return func(opt *optionsTemplate) {
		if opt.addFuncs == nil {
			opt.addFuncs = make(map[string]any)
		}

		k, v := fn(opt.opt)

		opt.addFuncs[k] = v
	}
}

func WithAddFuncMapWithOpts(fn func(o Option) map[string]any) OptionTemplate {
	return func(opt *optionsTemplate) {
		funcs := fn(opt.opt)

		if opt.addFuncs == nil {
			opt.addFuncs = make(map[string]any, len(funcs))
		}

		maps.Copy(opt.addFuncs, funcs)
	}
}
