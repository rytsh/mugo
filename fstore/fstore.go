package fstore

import (
	"io"
)

type ExecuteTemplate interface {
	ExecuteTemplate(wr io.Writer, name string, data any) error
}

func FuncMap(opts ...OptionFunc) map[string]any {
	opt := optionRun(opts...)

	return getFunc(opt)
}

func optionRun(opts ...OptionFunc) option {
	opt := option{
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

func getFunc(o option) map[string]any {
	v := valuer{
		Opt:   o,
		Value: make(map[string]any),
	}

	optReg := Option{
		Template: o.executeTemplate,
		WorkDir:  o.workDir,
		Trust:    o.trust,
		Log:      o.log,
	}

	for _, g := range reg.Groups {
		v.addGroup(g.name, g.fn)
	}

	for _, g := range reg.GroupsWithOptions {
		name, fn := g(optReg)
		v.addGroup(name, fn)
	}

	for name, fn := range reg.Funcs {
		v.addFunc(name, fn)
	}

	for _, f := range reg.FuncsWithOptions {
		name, fn := f(optReg)
		v.addFunc(name, fn)
	}

	return v.Value
}
