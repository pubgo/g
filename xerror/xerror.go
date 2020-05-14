package xerror

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"strings"
)

// NewXErr
func New(name string) xErr {
	name = replace(replace(strings.Title(name), " ", ""), "-", "")
	return xErr{name: name}
}

// XErr struct
type xErr struct {
	name string
}

func (t xErr) Error() string {
	return t.name
}

// Err err warp
func (t xErr) Err(format string, args ...interface{}) error {
	return fmt.Errorf(t.name+": "+format, args...)
}

// Msg err warp
func (t xErr) Msg(format string, args ...interface{}) string {
	return fmt.Sprintf(t.name+": "+format, args...)
}

// New
func (t xErr) New(format string) xErr {
	t.name = t.name + ":" + replace(replace(strings.Title(format), " ", ""), "-", "")
	return t
}

type xError struct {
	err     *error
	isPanic uint32
}

func WithErr(err *error) xError {
	return xError{err: err}
}

func (t *xError) Recover() {
	if t.isPanic == 0 {
		return
	}

	err := recover()
	if err == nil {
		return
	}

	if err1, ok := err.(error); ok {
		if err1 == ErrDone {
			return
		}

		*t.err = err1
		return
	}

	*t.err = fmt.Errorf("%#v", err)
}

func (t *xError) Panic(err error) {
	if isErrNil(err) {
		return
	}

	t.isPanic = 1
	panic(handle(err, ""))
}

func (t *xError) PanicF(err error, msg string, args ...interface{}) {
	if isErrNil(err) {
		return
	}

	t.isPanic = 1
	panic(handle(err, msg, args...))
}

func (t xError) Wrap(err error) error {
	if isErrNil(err) {
		return nil
	}

	return handle(err, "")
}

func (t xError) WrapF(err error) error {
	if isErrNil(err) {
		return nil
	}

	return handle(err, "")
}

// PanicErr
func (t *xError) PanicErr(d1 interface{}, err error) interface{} {
	if isErrNil(err) {
		return nil
	}

	t.isPanic = 1
	panic(handle(err, ""))
	return d1
}

// ExitErr
func (t xError) ExitErr(_ interface{}, err error) {
	if isErrNil(err) {
		return
	}

	fmt.Println(handle(err, "").Error())
	os.Exit(1)
}

// ExitF
func (t xError) ExitF(err error, msg string, args ...interface{}) {
	if isErrNil(err) {
		return
	}

	fmt.Println(handle(err, msg, args...).Error())
	os.Exit(1)
}

func (t xError) Exit(err error) {
	if isErrNil(err) {
		return
	}

	fmt.Println(handle(err, "").Error())
	os.Exit(1)
}

func UnWrap(err error) *xerror {
	err1, ok := err.(*xerror)
	if ok {
		return err1
	}

	return nil
}

type xerror struct {
	err    error
	Xrr    string  `json:"err,omitempty"`
	Msg    string  `json:"msg,omitempty"`
	Caller string  `json:"caller,omitempty"`
	Sub    *xerror `json:"sub,omitempty"`
}

func (t *xerror) P() {
	if t == nil {
		log.Println("xerror is nil")
		debug.PrintStack()
		os.Exit(1)
	}

	t.Xrr = t.err.Error()
	dt, _ := json.MarshalIndent(t, "", "\t")
	fmt.Println(string(dt))
}

func (t *xerror) next(e *xerror) {
	if t.Sub == nil {
		t.Sub = e
		return
	}
	t.Sub.next(e)
}

func (t *xerror) Is(err error) bool {
	if t == nil || t.err == nil || err == nil {
		return false
	}

	return t.err == err
}

func (t *xerror) As(err error) bool {
	if t == nil || t.err == nil || err == nil {
		return false
	}

	return strings.HasPrefix(t.err.Error(), err.Error())
}

func (t *xerror) Err() error {
	return t.err
}

// Error
func (t *xerror) Error() string {
	if t == nil || t.err == nil || t.err == ErrDone {
		return ""
	}

	t.Xrr = t.err.Error()
	dt, _ := json.Marshal(t)
	return string(dt)
}

func (t xerror) Reset() {
	t.err = nil
	t.Caller = ""
	t.Msg = ""
	if t.Sub != nil {
		t.Sub.Reset()
	}
}
