package registry

import (
	"reflect"

	"github.com/rytsh/call"
)

var (
	CallReg = call.NewReg()
	Group   = make(map[string]interface{})
)

func AddFunction(name string, fn interface{}, args ...string) {
	CallReg.AddFunction(name, fn, args...)
}

func AddGroup(name string, fn interface{}, args ...string) {
	CallReg.AddFunction(name, fn, args...)
	Group[name] = fn
}

func IsGroup(name string) bool {
	_, ok := Group[name]

	return ok
}

func ReturnWithFn(fn any) func() any {
	return func() any {
		return fn
	}
}

func ReturnWithPanic(v []any, err error) any {
	if err != nil {
		panic(err)
	}

	return v[0]
}

// GetFunc returns a function from the registry.
func GetFunc(fName string) any {
	v := ReturnWithPanic(CallReg.Call(fName))

	if _, ok := v.(map[string]interface{}); ok {
		return v
	}

	// check if v is a function
	if reflect.TypeOf(v).Kind() != reflect.Func {
		return ReturnWithFn(v)
	}

	return v
}
