package xerror

import (
	"github.com/pubgo/g/pkg"
	"github.com/pubgo/g/xenv"
	"github.com/pubgo/g/xinit"
)

// func caller depth
const callDepth = 2

var _isNone = func(val interface{}) bool {
	return val == nil || val == ErrDone || pkg.IsNone(val)
}
var _isZero = pkg.IsZero
var _caller = pkg.Caller
var skipErrorFile = xenv.IsSkipXerror()
var debug = xenv.IsDebug()

func init() {
	xinit.Init(func() (err error) {
		debug = xenv.IsDebug()
		skipErrorFile = xenv.IsSkipXerror()
		return
	})
}
