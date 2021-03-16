package recovery

import (
	"fmt"
	"os"
	"runtime/debug"
)

func Exit(args ...interface{}) {
	if r := recover(); r != nil {
		args = append(args, "\n", r, "\n")
		fmt.Print(args...)
		debug.PrintStack()
		os.Exit(1)
	}
}

func Check(fn func(err error)) {
	if r := recover(); r != nil {
		switch r := r.(type) {
		case error:
			fn(r)
		default:
			fn(fmt.Errorf("%#v\n", r))
		}
	}
}

func CheckErr(err *error) {
	if r := recover(); r != nil {
		switch r := r.(type) {
		case error:
			*err = r
		default:
			*err = fmt.Errorf("%#v\n", r)
		}
	}
}
