package xtry

import (
	"github.com/pubgo/xerror"
	"reflect"
	"sync"
)

func _TryRaw(fn reflect.Value) func(...reflect.Value) func(...reflect.Value) (err error) {
	if !assertFn(fn) {
		xerror.Exit(ErrNotFuncType)
	}

	var variadicType reflect.Value
	var isVariadic = fn.Type().IsVariadic()
	if isVariadic {
		variadicType = reflect.New(fn.Type().In(fn.Type().NumIn() - 1).Elem()).Elem()
	}

	numIn := fn.Type().NumIn()
	numOut := fn.Type().NumOut()
	return func(args ...reflect.Value) func(...reflect.Value) (err error) {
		if isVariadic && len(args) < numIn-1 {
			xerror.ExitF(ErrParamsNotMatch, "func %s input params error,func(%d,%d)", fn.Type(), numIn, len(args))
		}

		if !isVariadic && numIn != len(args) {
			xerror.ExitF(ErrParamsNotMatch, "func %s input params not match,func(%d,%d)", fn.Type(), numIn, len(args))
		}

		for i, k := range args {
			if !k.IsValid() || k.IsZero() {
				args[i] = reflect.New(fn.Type().In(i)).Elem()
				continue
			}

			if isVariadic {
				args[i] = variadicType
			}

			args[i] = k
		}

		return func(cfn ...reflect.Value) (err error) {
			defer xerror.RespErr(&err)
			defer valuePut(args)

			_c := fn.Call(args)
			if len(cfn) > 0 && cfn[0].IsValid() && !cfn[0].IsZero() {
				if cfn[0].Type().NumIn() != numOut {
					xerror.PanicF(ErrParamsNotMatch, "callback func input num and output num not match[%d]<->[%d]", cfn[0].Type().NumIn(), fn.Type().NumOut())
				}

				if cfn[0].Type().NumIn() != 0 && cfn[0].Type().In(0) != fn.Type().Out(0) {
					xerror.PanicF(ErrParamsNotMatch, "callback func out type error [%s]<->[%s]", cfn[0].Type().In(0), fn.Type().Out(0))
				}
				cfn[0].Call(_c)
			}
			return
		}
	}
}

func Try(fn func() error) (err error) {
	defer xerror.RespErr(&err)
	err = fn()

	return
}

func Pipe(fns ...func() error) (err error) {
	defer xerror.RespErr(&err)
	for _, fn := range fns {
		if err = fn(); err != nil {
			return
		}
	}
	return
}

// Wrap
func Wrap(fn interface{}) func(...interface{}) func(...interface{}) (err error) {
	_tr := _TryRaw(reflect.ValueOf(fn))
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
