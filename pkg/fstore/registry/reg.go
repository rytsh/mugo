package registry

var DirectReg = DReg{}

type DReg map[string]map[string]any

func (d DReg) Add(name string, v map[string]any) {
	d[name] = v
}
