package xerror

import (
	"encoding/json"
	"fmt"
)

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
	t.Caller = callerWithDepth()
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
		putXerror(t)
		return
	}

	sub := t.Sub
	t.Sub = nil
	putXerror(t)
	sub.Reset()
}
