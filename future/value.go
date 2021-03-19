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

func newFutureValue() *futureValue { return &futureValue{valChan: make(chan []reflect.Value, 1)} }

type futureValue struct {
	done    sync.Once
	values  []reflect.Value
	err     error
	valChan chan []reflect.Value
}

func (v *futureValue) Expect(format string, a ...interface{}) { xerror.PanicF(v.Err(), format, a...) }
func (v *futureValue) Err() error                             { _ = v.getVal(); return v.err }
func (v *futureValue) setErr(err error) *futureValue          { v.err = err; return v }
func (v *futureValue) Raw() []reflect.Value                   { return v.getVal() }
func (v *futureValue) String() string                         { return valueStr(v.getVal()...) }

func (v *futureValue) getVal() []reflect.Value {
	v.done.Do(func() {
		if v.valChan != nil {
			defer close(v.valChan)
			v.values = <-v.valChan
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
