package render

import (
	"github.com/Masterminds/sprig/v3"
	"github.com/rytsh/liz/utils/templatex"
	"github.com/rytsh/liz/utils/templatex/store"
	"github.com/rytsh/mugo/internal/render/generic"

	_ "github.com/rytsh/mugo/internal/render/funcs"
)

func FuncMap(trust bool, workDir string) func(t *templatex.Template) map[string]interface{} {
	return func(t *templatex.Template) map[string]interface{} {
		storeHolder := store.Holder{}
		storeHolder.AddFuncs(sprig.GenericFuncMap())

		// custom functions
		generic.CallReg.
			AddArgument("trust", trust).
			AddArgument("t", t).
			AddArgument("workDir", workDir)

		for _, fName := range generic.CallReg.GetFunctionNames() {
			storeHolder.AddFunc(fName, generic.GetFunc(fName))
		}

		return storeHolder.Funcs()
	}
}
