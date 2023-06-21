package store

import (
	"sync"
)

// Holder holds the functions.
type Holder struct {
	funcMap map[string]interface{}
	mutex   sync.RWMutex
}

func New(opts ...Option) *Holder {
	return new(Holder).initializeFuncs(opts...)
}

func (h *Holder) setFuncMap() {
	if h.funcMap == nil {
		h.funcMap = make(map[string]interface{})
	}
}

func (h *Holder) initializeFuncs(opts ...Option) *Holder {
	h.setFuncMap()

	option := &options{}
	for _, opt := range opts {
		opt(option)
	}

	h.AddFuncs(
		// add additional functions here
		option.addFuncs,
	)

	for _, f := range option.disableFuncs {
		delete(h.funcMap, f)
	}

	return h
}

// Funcs returns the functions.
func (h *Holder) Funcs() map[string]interface{} {
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	h.setFuncMap()

	return h.funcMap
}

// AddFuncs for adding multiple functions.
func (h *Holder) AddFuncs(funcs ...map[string]interface{}) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	h.setFuncMap()

	for _, f := range funcs {
		for k, v := range f {
			h.funcMap[k] = v
		}
	}
}

// AddFunc for adding a single function.
func (h *Holder) AddFunc(name string, fn interface{}) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	h.setFuncMap()

	h.funcMap[name] = fn
}
