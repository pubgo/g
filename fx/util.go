package fx

import (
	"context"
	"reflect"
	"runtime"
)

func CancelCtx() (ctx context.Context, cancel context.CancelFunc) {
	return context.WithCancel(context.Background())
}

// FnName ...
func FnName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

// ObjectName ...
func ObjectName(i interface{}) string {
	typ := reflect.TypeOf(i)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	return typ.PkgPath() + "." + typ.Name()
}

// CallerName ...
func CallerName(skip int) string {
	pc, _, _, _ := runtime.Caller(skip)
	return runtime.FuncForPC(pc).Name()
}
