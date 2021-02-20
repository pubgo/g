package xtest

import (
	"github.com/pubgo/x/pkg"
	"github.com/pubgo/x/pkg/colorutil"
	"github.com/pubgo/x/xtry"
)

var _caller = pkg.Caller
var _try = xtry.Try
var _isNone = pkg.IsNone

func init() {
	colorutil.Enable()
}
