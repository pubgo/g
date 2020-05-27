package xerror

import (
	"errors"
	"fmt"
	"runtime"
	"strconv"
)

func handleErr(err *error, _err interface{}) {
	if _err == nil {
		return
	}

	switch _err.(type) {
	case *xerror:
		*err = _err.(*xerror)
	case error:
		*err = _err.(error)
	case string:
		*err = errors.New(_err.(string))
	default:
		logger.Fatalf("unknown type, %#v", _err)
	}
}

func handle(err error, msg string, args ...interface{}) error {
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}

	err2 := xerrorPool.Get().(*xerror)
	err2.Msg = msg
	err2.Caller = callerWithDepth(callDepth)
	if err1, ok := err.(*xerror); ok {
		err2.Sub = err1
		err2.xrr = err1.xrr
		err1.xrr = nil

		err2.code = err1.code
		err1.code = 0
	} else {
		err2.xrr = err
	}

	return err2
}

func callerWithDepth(callDepth int) string {
	_, file, line, ok := runtime.Caller(callDepth)
	if !ok {
		return "no func caller error"
	}

	return file + ":" + strconv.Itoa(line)
}
