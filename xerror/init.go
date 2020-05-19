package xerror

import (
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

// func caller depth
const callDepth = 3

var (
	// ErrDone done
	ErrDone  = errors.New("DONE")
	isErrNil = func(err error) bool {
		return err == nil || err == ErrDone
	}
	replace = strings.ReplaceAll
	Debug   bool
	logger  = log.New(os.Stdout, "[xerror] ", log.LstdFlags|log.Lshortfile)
)

func init() {
	Debug = true

	if os.Getenv("debug") == "false" {
		Debug = false
	}
}

func handle(err error, msg string, args ...interface{}) error {
	if len(args) != 0 {
		msg = fmt.Sprintf(msg, args...)
	}

	if err1, ok := err.(*xerror); ok {
		err1.next(&xerror{Caller: callerWithDepth(callDepth), Msg: msg})
		return err1
	}

	return &xerror{err: err, Caller: callerWithDepth(callDepth), Msg: msg}
}

func callerWithDepth(callDepth int) string {
	_, file, line, ok := runtime.Caller(callDepth)
	if !ok {
		return "no func caller error"
	}
	return fmt.Sprintf("%s:%d", file, line)
}
