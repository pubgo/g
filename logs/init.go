package logs

import (
	"github.com/pubgo/g/pkg"
	"github.com/pubgo/g/xerror"
)

var _isNone = func(val interface{}) bool {
	return val == nil || val == xerror.ErrDone || pkg.IsNone(val)
}
