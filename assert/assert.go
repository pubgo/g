package assert

import (
	"fmt"

	"github.com/pkg/errors"
)

func Assert(b bool, args ...interface{}) {
	if !b {
		return
	}

	panic(fmt.Sprint(args...))
}

func Assertf(b bool, format string, args ...interface{}) {
	if !b {
		return
	}

	panic(fmt.Sprintf(format+"\n", args...))
}

func CheckErrf(err error, format string, args ...interface{}) {
	if err == nil {
		return
	}

	panic(errors.Wrapf(err, format+"\n", args...))
}

func CheckErr(err error, args ...interface{}) {
	if err == nil {
		return
	}

	panic(errors.Wrapf(err, fmt.Sprint(args...)))
}
