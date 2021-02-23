package xutil

import (
	"reflect"
	"runtime"
	"strconv"
	"strings"

	"github.com/pubgo/x/typeutil"
	"github.com/pubgo/xerror"
)

type frame uintptr

func (f frame) pc() uintptr { return uintptr(f) - 1 }

// CallerStack 调用栈
func CallerStack(cd int, fns ...func(fn *runtime.Func, pc uintptr) string) string {
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
	return strings.Join(typeutil.StrOf(file, ":", strconv.Itoa(line)), "")
}

// FuncStack 函数栈
func FuncStack(fn interface{}, fns ...func(fn *runtime.Func, pc uintptr) string) string {
	xerror.Assert(fn == nil, "[fn] is nil")

	var vfn = reflect.ValueOf(fn)
	xerror.Assert(!vfn.IsValid() || vfn.Kind() != reflect.Func || vfn.IsNil(), "[fn] func is nil or type error")

	var fnStack = runtime.FuncForPC(vfn.Pointer())
	if len(fns) > 0 {
		fns[0](fnStack, vfn.Pointer())
	}

	var file, line = fnStack.FileLine(vfn.Pointer())
	ma := strings.Split(fnStack.Name(), ".")
	return strings.Join(typeutil.StrOf(file, ":", strconv.Itoa(line), " ", ma[len(ma)-1]), "")
}
