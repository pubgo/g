package xerror

import (
	"github.com/pubgo/g/pkg"
	"github.com/pubgo/g/xenv"
	"reflect"
	"time"
)

// func caller depth
const callDepth = 2

var isErrNil = func(val error) bool {
	return val == nil || val == ErrDone
}
var _isZero = pkg.IsZero
var _caller = pkg.Caller
var skipErrorFile = xenv.IsSkipXerror()
var debug = true

func init() {
	debug = xenv.IsDebug()
	skipErrorFile = xenv.IsSkipXerror()
}

// IsZero exported
func IsNone(val reflect.Value) bool {
	if !val.IsValid() {
		return true
	}

	switch val.Kind() {
	case reflect.String:
		return val.Len() == 0
	case reflect.Bool:
		return val.Bool() == false
	case reflect.Float32, reflect.Float64:
		return val.Float() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return val.Uint() == 0
	case reflect.Ptr, reflect.Chan, reflect.Func, reflect.Interface, reflect.Slice, reflect.Map:
		return val.IsNil()
	case reflect.Array:
		for i := 0; i < val.Len(); i++ {
			if !IsNone(val.Index(i)) {
				return false
			}
		}
		return true
	case reflect.Struct:
		if t, ok := val.Interface().(time.Time); ok {
			return t.IsZero()
		}

		valid := val.FieldByName("Valid")
		if valid.IsValid() {
			va, ok := valid.Interface().(bool)
			return ok && !va
		}

		return reflect.DeepEqual(val.Interface(), reflect.Zero(val.Type()).Interface())
	default:
		return reflect.DeepEqual(val.Interface(), reflect.Zero(val.Type()).Interface())
	}
}
