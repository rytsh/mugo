package fstore

var reg = registry{
	Funcs: make(map[string]any),
}

type Option struct {
	Template ExecuteTemplate
	WorkDir  string
	Trust    bool
	Log      Adapter
}

type group struct {
	name string
	fn   func() map[string]any
}

type registry struct {
	Funcs             map[string]any
	FuncsWithOptions  []func(o Option) (string, any)
	Groups            []group
	GroupsWithOptions []func(o Option) (string, func() map[string]any)
}

func (r *registry) AddFunc(name string, fn any) {
	r.Funcs[name] = fn
}

func (r *registry) AddFuncWithOptions(fn func(o Option) (string, any)) {
	r.FuncsWithOptions = append(r.FuncsWithOptions, fn)
}

func (r *registry) AddGroup(g group) {
	r.Groups = append(r.Groups, g)
}

func (r *registry) AddGroupWithOptions(fn func(o Option) (string, func() map[string]any)) {
	r.GroupsWithOptions = append(r.GroupsWithOptions, fn)
}

func AddStruct(name string, s any) {
	reg.AddFunc(name, returnWithFn(s))
}

func AddStructWithOptions(fn func(o Option) (string, any)) {
	wrapfn := func(o Option) (string, any) {
		name, s := fn(o)
		return name, returnWithFn(s)
	}

	reg.AddFuncWithOptions(wrapfn)
}

func AddFunc(name string, fn any) {
	reg.AddFunc(name, fn)
}

func AddFuncWithOptions(fn func(o Option) (string, any)) {
	reg.AddFuncWithOptions(fn)
}

func AddGroup(name string, fn func() map[string]any) {
	reg.AddGroup(group{
		name: name,
		fn:   fn,
	})
}

func AddGroupWithOptions(fn func(o Option) (string, func() map[string]any)) {
	reg.AddGroupWithOptions(fn)
}

func GetRegistry() registry {
	return reg
}

// //////////////////////////////////////////////////

type Adapter interface {
	Error(msg string, keysAndValues ...any)
	Info(msg string, keysAndValues ...any)
	Debug(msg string, keysAndValues ...any)
	Warn(msg string, keysAndValues ...any)
}

type Noop struct{}

func (Noop) Error(_ string, _ ...any) {}
func (Noop) Info(_ string, _ ...any)  {}
func (Noop) Debug(_ string, _ ...any) {}
func (Noop) Warn(_ string, _ ...any)  {}
