package future

import (
	"fmt"
	"reflect"

	"github.com/pubgo/x/abc"
	"github.com/pubgo/x/fx"
	"github.com/pubgo/x/stack"
	"github.com/pubgo/xerror"
	"github.com/reactivex/rxgo/v2"
)

func Promise(provider rxgo.Producer, opts ...rxgo.Option) rxgo.Observable {
	return rxgo.Create([]rxgo.Producer{provider}, opts...)
}

func Async(fn interface{}, args ...interface{}) (val1 abc.FutureValue) {
	var value = newFutureValue()

	defer xerror.Resp(func(err xerror.XErr) { val1 = value.setErr(err) })

	xerror.Assert(fn == nil, "[fn] should not be nil")

	var vfn = fx.WrapRaw(fn)
	go func() {
		defer xerror.Resp(func(err1 xerror.XErr) {
			value.setErr(err1.WrapF("recovery error, input:%#v, func:%s, caller:%s",
				args, reflect.TypeOf(fn), stack.Func(fn)))
			value.valChan = nil
		})

		value.valChan <- vfn(args...)
	}()

	return value
}

func Await(val abc.FutureValue, fn interface{}, errs ...func(err error)) {
	err := val.Value(fn)
	if len(errs) > 0 {
		errs[0](err)
	}
}

func APipe(val abc.FutureValue, fn interface{}) (val1 abc.FutureValue) {
	var value = newFutureValue()

	defer xerror.Resp(func(err xerror.XErr) { val1 = value.setErr(err) })

	xerror.Assert(val == nil, "[val] should not be nil")
	xerror.Assert(fn == nil, "[fn] should not be nil")

	var vfn = fx.WrapReflect(fn)
	go func() {
		if err := val.Err(); err != nil {
			value.setErr(err)
			value.valChan = nil
			return
		}

		defer xerror.Resp(func(err1 xerror.XErr) {
			value.setErr(err1.WrapF("input:%s, func:%s", val.Raw(), reflect.TypeOf(fn)))
			value.valChan <- make([]reflect.Value, 0)
		})

		value.valChan <- vfn(val.Raw()...)
	}()

	return value
}

func valueStr(values ...reflect.Value) string {
	var data []interface{}
	for _, dt := range values {
		var val interface{}
		if dt.IsValid() {
			val = dt.Interface()
		}
		data = append(data, val)
	}
	return fmt.Sprint(data...)
}

func RunComplete(values ...abc.FutureValue) (err error) {
	defer xerror.RespErr(&err)

	for i := range values {
		err = xerror.Append(err, values[i].Err())
	}

	return xerror.Wrap(err)
}
