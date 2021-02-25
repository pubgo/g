package future

import (
	"reflect"
	"sync"

	"github.com/pubgo/x/abc"
	"github.com/pubgo/x/fx"
	"github.com/pubgo/xerror"
	"github.com/pubgo/xerror/xerror_abc"
)

var _ abc.FutureValue = (*futureValue)(nil)

func newFutureValue() *futureValue { return &futureValue{} }

type futureValue struct {
	done   sync.Once
	values []reflect.Value
	err    error
	valFn  func() []reflect.Value
}

func (v *futureValue) Expect(format string, a ...interface{}) { xerror.PanicF(v.Err(), format, a...) }
func (v *futureValue) Err() error                             { _ = v.getVal(); return v.err }
func (v *futureValue) setErr(err error) *futureValue          { v.err = err; return v }
func (v *futureValue) Raw() []reflect.Value                   { return v.getVal() }
func (v *futureValue) String() string                         { return valueStr(v.getVal()...) }

func (v *futureValue) Get() interface{} {
	val := v.getVal()
	if len(val) == 0 || !val[0].IsValid() {
		return nil
	}

	return val[0].Interface()
}

func (v *futureValue) getVal() []reflect.Value {
	v.done.Do(func() {
		if v.valFn != nil {
			v.values = v.valFn()
		}
	})
	return v.values
}

func (v *futureValue) Value(fn interface{}) (gErr error) {
	defer xerror.Resp(func(err xerror_abc.XErr) {
		gErr = err.WrapF("input:%s, func:%s", valueStr(v.getVal()...), reflect.TypeOf(fn))
	})

	xerror.Assert(fn == nil, "[fn] should not be nil")
	xerror.Panic(v.Err())

	fx.WrapReflect(fn)(v.getVal()...)
	return
}

var _ abc.Value = (*value)(nil)

type value struct {
	err error
	val interface{}
}

func (v *value) Assert(format string, a ...interface{}) { xerror.PanicF(v.Err(), format, a...) }
func (v *value) Err() error                             { return v.err }
func (v *value) Value() interface{}                     { return v.val }
