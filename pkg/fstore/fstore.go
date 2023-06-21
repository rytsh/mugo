package fstore

import (
	"github.com/rytsh/mugo/pkg/fstore/funcs"
	"github.com/rytsh/mugo/pkg/fstore/registry"
)

func FuncMap(opts ...Option) map[string]interface{} {
	opt := optionRun(opts...)

	return funcX(opt)(opt.executeTemplate)
}

func FuncMapTpl(opts ...Option) func(t funcs.ExecuteTemplate) map[string]interface{} {
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

func funcX(o options) func(t funcs.ExecuteTemplate) map[string]interface{} {
	return func(t funcs.ExecuteTemplate) map[string]interface{} {
		v := make(map[string]interface{})

		// custom functions
		registry.CallReg.
			AddArgument("trust", o.trust).
			AddArgument("log", nil).
			AddArgument("template", t).
			AddArgument("workDir", o.workDir)

		for gName, mapValues := range registry.DirectReg {
			if _, ok := o.disableGroups[gName]; ok {
				continue
			}

			if !isSpecificGroup(o, gName) {
				continue
			}

			for fName, fValue := range mapValues {
				if _, ok := o.disableFuncs[fName]; ok {
					continue
				}

				if !isSpecificFunc(o, fName) {
					continue
				}

				v[fName] = fValue
			}
		}

		if len(o.specificGroups) == 0 {
			for _, fName := range registry.CallReg.GetFunctionNames() {
				if _, ok := o.disableFuncs[fName]; ok {
					continue
				}

				if !isSpecificFunc(o, fName) {
					continue
				}

				v[fName] = registry.GetFunc(fName)
			}
		}

		return v
	}
}

func isSpecificFunc(o options, name string) bool {
	if len(o.specificFunc) > 0 {
		if _, ok := o.specificFunc[name]; ok {
			return true
		}

		return false
	}

	return true
}

func isSpecificGroup(o options, name string) bool {
	if len(o.specificGroups) > 0 {
		if _, ok := o.specificGroups[name]; ok {
			return true
		}

		return false
	}

	return true
}
