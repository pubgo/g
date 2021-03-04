package xutil

import (
	"github.com/pubgo/x/fx"
	"github.com/pubgo/xerror"

	"reflect"
	"unsafe"
)

// #nosec G103
// ToStr returns a string pointer without allocation
func ToStr(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

// #nosec G103
// ToBytes returns a byte pointer without allocation
func ToBytes(str string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&str))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// If exported
func If(check bool, a, b interface{}) interface{} {
	if check {
		return a
	}
	return b
}

func Default(t, f interface{}) interface{} {
	if IsZero(t) {
		return f
	}
	return t
}

// IsZero
func IsZero(val interface{}) bool {
	return val == nil || reflect.ValueOf(val).IsZero()
}

func TryCatch(fn func(), catch ...func(err xerror.XErr)) {
	xerror.Assert(fn == nil, "[fn] should not be nil")

	defer xerror.Resp(func(err xerror.XErr) {
		if len(catch) > 0 {
			catch[0](err)
		}
	})

	fn()
}

func TryWith(err *error, fn interface{}, args ...interface{}) {
	xerror.Assert(fn == nil, "[fn] should not be nil")

	defer xerror.RespErr(err)
	fx.WrapValue(fn, args...)

	return
}

func Try(fn interface{}, args ...interface{}) (gErr xerror.XErr) {
	xerror.Assert(fn == nil, "[fn] should not be nil")

	defer xerror.Resp(func(err xerror.XErr) { gErr = err })
	fx.WrapValue(fn, args...)

	return
}
