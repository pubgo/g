package xerror

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

// func caller depth
const callDepth = 3

var (
	// ErrDone done
	ErrDone  = errors.New("DONE")
	isErrNil = func(val error) bool {
		return val == nil || val == ErrDone
	}
	replace = strings.ReplaceAll
)

func handle(err error, msg string, args ...interface{}) error {
	if err1, ok := err.(*xerror); ok {
		err1.next(&xerror{Caller: callerWithDepth(callDepth), Msg: fmt.Sprintf(msg, args...)})
		return err1
	}

	return &xerror{err: err, Caller: callerWithDepth(callDepth), Msg: fmt.Sprintf(msg, args...)}
}

func callerWithDepth(callDepth int) string {
	_, file, line, ok := runtime.Caller(callDepth)
	if !ok {
		return "no func caller error"
	}
	return fmt.Sprintf("%s:%d", file, line)
}
