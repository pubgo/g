package xtry

import (
	"reflect"
)

func assertFn(fn reflect.Value) bool {
	if !fn.IsValid() || fn.IsNil() || fn.Kind() != reflect.Func {
		return false
	}
	return true
}
