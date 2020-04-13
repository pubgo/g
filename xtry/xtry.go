package xtry

import (
	"github.com/pubgo/g/xerror"
	"reflect"
	"sync"
)

func _TryRaw(fn reflect.Value) func(...reflect.Value) func(...reflect.Value) (err error) {
	xerror.PanicM(xerror.AssertFn(fn), ErrNotFuncType)

	var variadicType reflect.Value
	var isVariadic = fn.Type().IsVariadic()
	if isVariadic {
		variadicType = reflect.New(fn.Type().In(fn.Type().NumIn() - 1).Elem()).Elem()
	}

	_NumIn := fn.Type().NumIn()
	return func(args ...reflect.Value) func(...reflect.Value) (err error) {
		if isVariadic && len(args) < _NumIn-1 {
			xerror.PanicM(ErrParamsNotMatch, "func %s input params error,func(%d,%d)", fn.Type(), _NumIn, len(args))
		}

		if !isVariadic && _NumIn != len(args) {
			xerror.PanicM(ErrParamsNotMatch, "func %s input params not match,func(%d,%d)", fn.Type(), _NumIn, len(args))
		}

		for i, k := range args {
			if _isZero(k) {
				args[i] = reflect.New(fn.Type().In(i)).Elem()
				continue
			}

			if isVariadic {
				args[i] = variadicType
			}

			args[i] = k
		}

		return func(cfn ...reflect.Value) (err error) {
			defer func() {
				xerror.ErrHandle(recover(), func(_err xerror.IErr) {
					_fn := fn
					if len(cfn) > 0 && !_isZero(cfn[0]) {
						_fn = cfn[0]
					}
					_err.Caller(_caller.FromFunc(_fn))
					err = _err
				})
			}()

			_c := fn.Call(args)
			if len(cfn) > 0 && !_isZero(cfn[0]) {
				xerror.PanicM(xerror.AssertFn(cfn[0]), ErrNotFuncType)
				if cfn[0].Type().NumIn() != fn.Type().NumOut() {
					xerror.PanicM(ErrParamsNotMatch, "callback func input num and output num not match[%d]<->[%d]", cfn[0].Type().NumIn(), fn.Type().NumOut())
				}

				if cfn[0].Type().NumIn() != 0 && cfn[0].Type().In(0) != fn.Type().Out(0) {
					xerror.PanicM(ErrParamsNotMatch, "callback func out type error [%s]<->[%s]", cfn[0].Type().In(0), fn.Type().Out(0))
				}
				cfn[0].Call(_c)
			}
			return
		}
	}
}

// Try xerror
func Try(fn interface{}) func(...interface{}) func(...interface{}) (err error) {
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
