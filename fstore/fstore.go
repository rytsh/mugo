package fstore

import (
	"io"

	"github.com/rytsh/mugo/fstore/registry/cast"
	"github.com/rytsh/mugo/fstore/registry/codec"
	"github.com/rytsh/mugo/fstore/registry/crypto"
	"github.com/rytsh/mugo/fstore/registry/exec"
	"github.com/rytsh/mugo/fstore/registry/external"
	"github.com/rytsh/mugo/fstore/registry/faker"
	"github.com/rytsh/mugo/fstore/registry/file"
	"github.com/rytsh/mugo/fstore/registry/html2"
	"github.com/rytsh/mugo/fstore/registry/humanize"
	"github.com/rytsh/mugo/fstore/registry/log"
	"github.com/rytsh/mugo/fstore/registry/maps"
	"github.com/rytsh/mugo/fstore/registry/math"
	"github.com/rytsh/mugo/fstore/registry/minify"
	"github.com/rytsh/mugo/fstore/registry/os"
	"github.com/rytsh/mugo/fstore/registry/random"
	"github.com/rytsh/mugo/fstore/registry/template"
	"github.com/rytsh/mugo/fstore/registry/time"
	"github.com/rytsh/mugo/fstore/registry/ungroup"
)

type ExecuteTemplate interface {
	ExecuteTemplate(wr io.Writer, name string, data any) error
}

func FuncMap(opts ...Option) map[string]interface{} {
	opt := optionRun(opts...)

	return funcX(opt)(opt.executeTemplate)
}

func FuncMapTpl(opts ...Option) func(t ExecuteTemplate) map[string]interface{} {
	return funcX(optionRun(opts...))
}

func optionRun(opts ...Option) option {
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

func funcX(o option) func(t ExecuteTemplate) map[string]interface{} {
	return func(t ExecuteTemplate) map[string]interface{} {
		v := valuer{
			Opt:   o,
			Value: make(map[string]interface{}),
		}

		v.addGroup("sprig", external.Sprig)
		v.addGroup("ungroup", ungroup.Ungroup)

		v.addFunc("exec", exec.New(o.trust, o.log).Exec)
		v.addFunc("execTemplate", template.New(t).ExecTemplate)

		v.addFunc("cast", returnWithFn(cast.Cast{}))
		v.addFunc("codec", returnWithFn(codec.Codec{}))
		v.addFunc("crypto", returnWithFn(crypto.Crypto{}))
		v.addFunc("faker", returnWithFn(faker.New()))
		v.addFunc("file", returnWithFn(file.New(o.trust)))
		v.addFunc("html2", returnWithFn(html2.HTML2{}))
		v.addFunc("humanize", returnWithFn(humanize.Humanize{}))
		v.addFunc("log", returnWithFn(log.Log{}))
		v.addFunc("map", returnWithFn(maps.New()))
		v.addFunc("math", returnWithFn(math.Math{}))
		v.addFunc("minify", returnWithFn(minify.New()))
		v.addFunc("os", returnWithFn(os.New(o.workDir)))
		v.addFunc("random", returnWithFn(random.New(nil)))
		v.addFunc("time", returnWithFn(time.Time{}))

		return v.Value
	}
}
