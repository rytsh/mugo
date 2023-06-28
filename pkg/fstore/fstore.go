package fstore

import (
	_ "github.com/rytsh/mugo/pkg/fstore/funcs"
	"github.com/rytsh/mugo/pkg/fstore/registry"
)

func FuncMap(opts ...Option) map[string]interface{} {
	opt := optionRun(opts...)

	return funcX(opt)(opt.executeTemplate)
}

func FuncMapTpl(opts ...Option) func(t registry.ExecuteTemplate) map[string]interface{} {
	return funcX(optionRun(opts...))
}

func optionRun(opts ...Option) options {
	opt := options{
		disableFuncs:   make(map[string]struct{}),
		disableGroups:  make(map[string]struct{}),
		specificFunc:   make(map[string]struct{}),
		specificGroups: make(map[string]struct{}),
	}
	for _, o := range opts {
		o(&opt)
	}

	return opt
}

func funcX(o options) func(t registry.ExecuteTemplate) map[string]interface{} {
	return func(t registry.ExecuteTemplate) map[string]interface{} {
		v := make(map[string]interface{})

		// custom functions
		registry.CallReg.
			AddArgument("trust", o.trust).
			AddArgument("log", nil).
			AddArgument("template", t).
			AddArgument("workDir", o.workDir)

		for _, fName := range registry.CallReg.GetFunctionNames() {
			// fname is a group and not a specific group
			if registry.IsGroup(fName) && !isSpecificGroup(o, fName) {
				continue
			}

			switch vTyped := registry.GetFunc(fName).(type) {
			case map[string]interface{}:
				for key, value := range vTyped {
					if !isSpecificFunc(o, key) {
						continue
					}

					v[key] = value
				}
			default:
				if !isSpecificFunc(o, fName) {
					continue
				}

				v[fName] = vTyped
			}
		}

		return v
	}
}

func isSpecificFunc(o options, name string) bool {
	if _, ok := o.disableFuncs[name]; ok {
		return false
	}

	if len(o.specificFunc) > 0 {
		if _, ok := o.specificFunc[name]; ok {
			return true
		}

		return false
	}

	return true
}

func isSpecificGroup(o options, name string) bool {
	if _, ok := o.disableGroups[name]; ok {
		return false
	}

	if len(o.specificGroups) > 0 {
		if _, ok := o.specificGroups[name]; ok {
			return true
		}

		return false
	}

	return true
}
