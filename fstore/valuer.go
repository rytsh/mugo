package fstore

type valuer struct {
	Opt   option
	Value map[string]interface{}
}

func (s *valuer) addFunc(name string, fn any) {
	if !isEnable(name, s.Opt, false) {
		return
	}

	if fn == nil {
		return
	}

	s.Value[name] = fn
}

func (s *valuer) addGroup(name string, fn func() map[string]interface{}) {
	if !isEnable(name, s.Opt, true) {
		return
	}

	if fn == nil {
		return
	}

	for k, f := range fn() {
		if !isEnable(k, s.Opt, false) {
			continue
		}

		if f == nil {
			continue
		}

		s.Value[k] = f
	}
}

func isEnable(name string, o option, checkGroup bool) bool {
	if checkGroup {
		return isEnableGroup(name, o)
	}

	return isEnableFunc(name, o)
}

func isEnableGroup(name string, o option) bool {
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

func isEnableFunc(name string, o option) bool {
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

func returnWithFn[T any](fn T) func() T {
	return func() T {
		return fn
	}
}
