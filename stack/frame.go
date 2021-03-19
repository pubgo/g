package stack

import (
	"fmt"
	"reflect"
	"runtime"

	"github.com/pubgo/xerror"
)

type frame uintptr

func (f frame) pc() uintptr { return uintptr(f) - 1 }

// Caller 调用栈
func Caller(cd int, fns ...func(fn *runtime.Func, pc uintptr) string) string {
	var pcs = make([]uintptr, 1)
	if runtime.Callers(cd, pcs[:]) == 0 {
		return ""
	}

	f := frame(pcs[0])
	fn := runtime.FuncForPC(f.pc())
	if fn == nil {
		return "unknown type"
	}

	if len(fns) > 0 {
		return fns[0](fn, f.pc())
	}

	file, line := fn.FileLine(f.pc())
	return fmt.Sprintf("%s:%d", file, line)
}

// Func 函数栈
func Func(fn interface{}, fns ...func(fn *runtime.Func, pc uintptr) string) string {
	xerror.Assert(fn == nil, "[fn] is nil")

	var vfn = reflect.ValueOf(fn)
	var p = vfn.Pointer()
	xerror.Assert(!vfn.IsValid() || vfn.Kind() != reflect.Func || vfn.IsNil(), "[fn] func is nil or type error")

	var fnStack = runtime.FuncForPC(p)
	if len(fns) > 0 {
		return fns[0](fnStack, p)
	}

	var file, line = fnStack.FileLine(p)
	return fmt.Sprintf("%s:%d <%s>", file, line, vfn.Type().String())
}
