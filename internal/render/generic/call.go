package generic

import (
	"reflect"

	"github.com/rytsh/call"
)

var CallReg = call.NewReg()

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

func GetFunc(fName string) any {
	v := ReturnWithPanic(CallReg.Call(fName))

	// check if v is a function
	if reflect.TypeOf(v).Kind() != reflect.Func {
		return ReturnWithFn(v)
	}

	return v
}
