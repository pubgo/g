package future

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/pubgo/x/abc"
	"github.com/pubgo/x/fx"
	"github.com/pubgo/x/stack"
	"github.com/pubgo/x/sync2"
	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog"
	"go.uber.org/atomic"
)

type promise struct {
	stack  string
	closed atomic.Bool
	done   sync.Once
	err    atomic.Error
	wg     sync2.WaitGroup
	data   chan abc.FutureValue
}

func (s *promise) cancel()        { s.closed.Store(true) }
func (s *promise) waitForClose()  { s.done.Do(func() { go func() { s.wg.Wait(); close(s.data) }() }) }
func (s *promise) isClosed() bool { return s.closed.Load() }
func (s *promise) Await() chan abc.FutureValue {
	s.waitForClose()

	var val = make(chan abc.FutureValue)
	go func() {
		defer close(val)
		for data := range s.data {
			val <- data
			s.wg.Done()
		}
	}()

	return val
}

func (s *promise) yield(data interface{}) {
	if s.isClosed() {
		xlog.Warnf("[xprocess] [promise] %s closed", s.stack)
		return
	}

	switch val := data.(type) {
	case func():
		s.Go(val)
	case *futureValue:
		s.wg.Inc()
		go func() { s.data <- val }()
	default:
		s.wg.Inc()
		value := newFutureValue()
		value.values = []reflect.Value{reflect.ValueOf(data)}
		go func() { s.data <- value }()
	}
}

func (s *promise) await(val abc.FutureValue, fn interface{}) {
	if s.isClosed() {
		xlog.Warnf("[xprocess] [promise] %s closed", s.stack)
		return
	}

	s.yield(Await(val, fn))
}

func (s *promise) Map(fn interface{}) (val1 abc.Value) {
	s.waitForClose()

	defer xerror.Resp(func(err xerror.XErr) { val1 = &value{err: err} })

	xerror.Assert(fn == nil, "[fn] should not be nil")

	var errs []error
	var values []abc.FutureValue
	for data := range s.Await() {
		if err := data.Err(); err != nil {
			errs = append(errs, err)
		}

		values = append(values, Await(data, fn))
	}

	// 遇到未知错误
	xerror.PanicErrs(errs...)

	if len(values) == 0 {
		return
	}

	tfn := reflect.TypeOf(fn)

	// fn没有输出
	if tfn.NumOut() == 0 {
		return &value{}
	}

	var t = tfn.Out(0)
	var rst = reflect.MakeSlice(reflect.SliceOf(t), 0, len(values))
	for i := range values {
		val := values[i].Raw()[0]

		if !val.IsValid() {
			if t.Kind() == reflect.Ptr {
				val = reflect.New(t).Elem()
			} else {
				val = reflect.Zero(t)
			}
		}

		rst = reflect.Append(rst, val)
	}
	xerror.PanicErrs(errs...)

	return &value{val: rst.Interface()}
}

func (s *promise) Go(fn func()) {
	s.wg.Inc()
	go func() {
		defer s.wg.Done()
		defer xerror.Resp(func(err xerror.XErr) { s.cancel(); s.err.Store(err) })
		fn()
	}()
}

func Promise(fn func(g abc.Future)) abc.IPromise {
	xerror.Assert(fn == nil, "[fn] should not be nil")

	s := &promise{data: make(chan abc.FutureValue), stack: stack.Func(fn)}
	s.Go(func() { fn(&future{p: s}) })
	return s
}

func Async(fn interface{}, args ...interface{}) (val1 abc.FutureValue) {
	var value = newFutureValue()

	defer xerror.Resp(func(err xerror.XErr) { val1 = value.setErr(err) })

	xerror.Assert(fn == nil, "[fn] should not be nil")

	var vfn = fx.WrapRaw(fn)
	var data = make(chan []reflect.Value, 1)
	value.valFn = func() []reflect.Value { return <-data }
	go func() {
		defer xerror.Resp(func(err1 xerror.XErr) {
			value.setErr(err1.WrapF("recovery error, input:%#v, func:%s, caller:%s",
				args, reflect.TypeOf(fn), stack.Func(fn)))
			data <- make([]reflect.Value, 0)
		})

		data <- vfn(args...)
	}()

	return value
}

func Await(val abc.FutureValue, fn interface{}) (val1 abc.FutureValue) {
	var value = newFutureValue()

	defer xerror.Resp(func(err xerror.XErr) { val1 = value.setErr(err) })

	xerror.Assert(val == nil, "[val] should not be nil")
	xerror.Assert(fn == nil, "[fn] should not be nil")

	var data = make(chan []reflect.Value, 1)
	value.valFn = func() []reflect.Value { return <-data }

	var vfn = fx.WrapReflect(fn)
	go func() {
		if err := val.Err(); err != nil {
			value.setErr(err)
			data <- make([]reflect.Value, 0)
			return
		}

		defer xerror.Resp(func(err1 xerror.XErr) {
			value.setErr(err1.WrapF("input:%s, func:%s", val.Raw(), reflect.TypeOf(fn)))
			data <- make([]reflect.Value, 0)
		})

		data <- vfn(val.Raw()...)
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

	var errs []error
	for i := range values {
		if err := values[i].Err(); err != nil {
			errs = append(errs, err)
		}
	}

	xerror.PanicErrs(errs...)
	return nil
}
