package xerror

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"runtime/debug"
)

type XErr interface {
	error
	As(err interface{}) bool
	Is(err error) bool
	Unwrap() error
	Wrap(err error) error
	Code() int
}

func RespErr(err *error) {
	handleErr(err, recover())
}

func Resp(f func(err XErr)) {
	var err error
	handleErr(&err, recover())
	if err != nil {
		f(err.(XErr))
		err.(*xerror).Reset()
	}
}

func Panic(err error) {
	if isErrNil(err) {
		return
	}

	panic(handle(err, ""))
}

func PanicF(err error, msg string, args ...interface{}) {
	if isErrNil(err) {
		return
	}

	panic(handle(err, msg, args...))
}

func Wrap(err error) error {
	if isErrNil(err) {
		return nil
	}

	return handle(err, "")
}

func WrapF(err error, msg string, args ...interface{}) error {
	if isErrNil(err) {
		return nil
	}

	return handle(err, msg, args...)
}

// PanicErr
func PanicErr(d1 interface{}, err error) interface{} {
	if isErrNil(err) {
		return d1
	}

	panic(handle(err, ""))
}

// ExitErr
func ExitErr(_ interface{}, err error) {
	if isErrNil(err) {
		return
	}

	logger.Println(handle(err, "").Error())
	if Debug {
		debug.PrintStack()
	}
	os.Exit(1)
}

// ExitF
func ExitF(err error, msg string, args ...interface{}) {
	if isErrNil(err) {
		return
	}

	logger.Println(handle(err, msg, args...).Error())
	if Debug {
		debug.PrintStack()
	}
	os.Exit(1)
}

func Exit(err error) {
	if isErrNil(err) {
		return
	}

	logger.Println(handle(err, "").Error())
	if Debug {
		debug.PrintStack()
	}
	os.Exit(1)
}

// ext from errors
var (
	UnWrap = errors.Unwrap
	Is     = errors.Is
	As     = func(err error, target interface{}) bool {
		if target == nil {
			return false
		}

		val := reflect.ValueOf(target)
		typ := val.Type()

		if typ.Kind() != reflect.Ptr || val.IsNil() {
			return false
		}

		if e := typ.Elem(); e.Kind() != reflect.Interface && !typ.Implements(errorType) {
			return false
		}

		targetType := typ.Elem()
		for err != nil {
			if reflect.TypeOf(err).AssignableTo(targetType) {
				val.Elem().Set(reflect.ValueOf(err))
				return true
			}
			if x, ok := err.(interface{ As(interface{}) bool }); ok && x.As(target) {
				return true
			}
			err = UnWrap(err)
		}
		return false
	}
	errorType = reflect.TypeOf((*error)(nil)).Elem()
)

func New(code int, Msg string) interface {
	Wrap(error) error
	Code() int
} {
	return &xerror{code: code, Msg: Msg}
}

type xerror struct {
	error
	xrr    error
	code   int
	Err    string  `json:"err,omitempty"`
	Msg    string  `json:"msg,omitempty"`
	Caller string  `json:"caller,omitempty"`
	Sub    *xerror `json:"sub,omitempty"`
}

func (t *xerror) Code() int {
	return t.code
}

func (t *xerror) next(e *xerror) {
	if t.Sub == nil {
		t.Sub = e
		return
	}
	t.Sub.next(e)
}

func (t *xerror) Unwrap() error {
	if t == nil {
		return nil
	}

	return t.xrr
}

func (t *xerror) Wrap(err error) error {
	if isErrNil(err) {
		return nil
	}

	t.xrr = err
	t.Caller = callerWithDepth(callDepth)
	return t
}

func (t *xerror) Is(err error) bool {
	if t == nil {
		return false
	}

	return t.xrr == err
}

func (t *xerror) As(err interface{}) bool {
	if t == nil || t.xrr == nil || err == nil {
		return false
	}

	switch e := err.(type) {
	case *xerror:
		fmt.Println(e.code)
		return t.code == e.code
	case int:
		return t.code == e
	case string:
		return t.Msg == e
	default:
		return false
	}
}

// Error
func (t *xerror) Error() string {
	if t == nil || t.xrr == nil || t.xrr == ErrDone {
		return ""
	}

	t.Err = t.xrr.Error()
	var dt []byte

	if Debug {
		dt, _ = json.MarshalIndent(t, "", "\t")
	} else {
		dt, _ = json.Marshal(t)
	}
	return string(dt)
}

func (t *xerror) Reset() {
	t.xrr = nil
	t.code = 0
	t.Err = ""
	t.Msg = ""
	t.Caller = ""
	if t.Sub == nil {
		xerrorPool.Put(t)
		return
	}

	sub := t.Sub
	t.Sub = nil
	xerrorPool.Put(t)
	sub.Reset()
}
