package functions

import (
	"sync"
	textTemplate "text/template"

	"github.com/Masterminds/sprig/v3"

	"github.com/rytsh/mugo/pkg/template/functions/custom"
	"github.com/rytsh/mugo/pkg/template/functions/hugo"
	"github.com/rytsh/mugo/pkg/template/functions/humanize"
)

var Global = Holder{}

type Holder struct {
	funcMap    map[string]interface{}
	mutex      sync.RWMutex
	initialize sync.Once
}

func (h *Holder) InitializeFuncs() *Holder {
	h.initialize.Do(func() {
		if h.funcMap == nil {
			h.funcMap = make(map[string]interface{})
		}

		h.AddFuncs(
			sprig.GenericFuncMap(),
			hugo.FuncMap(),
			humanize.FuncMap(),
			custom.FuncMap(),
			// Add additonal functions here
		)
	})

	return h
}

func (h *Holder) Funcs() map[string]interface{} {
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	return h.funcMap
}

func (h *Holder) TxtFuncs() textTemplate.FuncMap {
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	return h.funcMap
}

func (h *Holder) AddFuncs(funcs ...map[string]interface{}) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	for _, f := range funcs {
		for k, v := range f {
			h.funcMap[k] = v
		}
	}
}
