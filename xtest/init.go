package xtest

import (
	"github.com/pubgo/g/pkg"
	"github.com/pubgo/g/pkg/colorutil"
	"github.com/pubgo/g/xtry"
)

var _caller = pkg.Caller
var _try = xtry.Try
var _isNone = pkg.IsNone

func init() {
	colorutil.Enable()
}
