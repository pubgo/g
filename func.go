package x

import (
	"fmt"
	"github.com/pubgo/xerror"
	"reflect"
	"sync"
)

func Func(fn interface{}) func(...interface{}) func(...interface{}) (err error) {
	if fn == nil {
		xerror.Panic(ErrParamIsNil)
	}

	_tr := tryWrap(reflect.ValueOf(fn))
	return func(args ...interface{}) func(...interface{}) (err error) {
		var _args = valueGet()
		defer valuePut(_args)

		for _, k := range args {
			_args = append(_args, reflect.ValueOf(k))
		}
		_tr1 := _tr(_args...)
		return func(cfn ...interface{}) (err error) {
			var _cfn = valueGet()
			defer valuePut(_cfn)

			for _, k := range cfn {
				_cfn = append(_cfn, reflect.ValueOf(k))
			}
			return _tr1(_cfn...)
		}
	}
}

func tryWrap(fn reflect.Value) func(...reflect.Value) func(...reflect.Value) (err error) {
	if fn.Type().Kind() != reflect.Func {
		xerror.Panic(ErrParamTypeNotFunc)
	}

	numIn := fn.Type().NumIn()
	var variadicType reflect.Value
	var isVariadic = fn.Type().IsVariadic()
	if isVariadic {
		variadicType = reflect.New(fn.Type().In(numIn - 1).Elem()).Elem()
	}

	numOut := fn.Type().NumOut()
	return func(args ...reflect.Value) func(...reflect.Value) (err error) {
		if !isVariadic && numIn != len(args) || isVariadic && len(args) < numIn-1 {
			xerror.PanicF(ErrInputParamsNotMatch, "func: %s, func(%d,%d)", fn.Type(), numIn, len(args))
		}

		for i, k := range args {
			if !k.IsValid() {
				args[i] = reflect.New(fn.Type().In(i)).Elem()
				continue
			}

			switch k.Kind() {
			case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice, reflect.UnsafePointer:
				if k.IsNil() {
					args[i] = reflect.New(fn.Type().In(i)).Elem()
					continue
				}
			}

			if isVariadic {
				args[i] = variadicType
			}

			args[i] = k
		}

		return func(cfn ...reflect.Value) (err error) {
			defer func() {
				if err1 := recover(); err1 != nil {
					switch err1 := err1.(type) {
					case error:
						err = err1
					default:
						err = fmt.Errorf("%s -> [func: %s] [input: %s]", err1, fn.String(), args)
					}
				}
			}()
			defer valuePut(args)

			_c := fn.Call(args)
			if len(cfn) > 0 && cfn[0].IsValid() && !cfn[0].IsZero() {
				if cfn[0].Type().NumIn() != numOut {
					xerror.PanicF(ErrInputOutputParamsNotMatch, "[%d]<->[%d]", cfn[0].Type().NumIn(), fn.Type().NumOut())
				}

				if cfn[0].Type().NumIn() != 0 && cfn[0].Type().In(0) != fn.Type().Out(0) {
					xerror.PanicF(ErrFuncOutputTypeNotMatch, "[%s]<->[%s]", cfn[0].Type().In(0), fn.Type().Out(0))
				}
				cfn[0].Call(_c)
			}
			return
		}
	}
}

var _valuePool = sync.Pool{
	New: func() interface{} {
		return []reflect.Value{}
	},
}

func valueGet() []reflect.Value {
	return _valuePool.Get().([]reflect.Value)
}

func valuePut(v []reflect.Value) {
	_valuePool.Put(v[:0])
}

func SliceOf(args ...interface{}) []interface{} {
	return args
}
