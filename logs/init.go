package logs

import (
	"github.com/pubgo/x/pkg"
	"github.com/pubgo/x/xerror"
)

var _isNone = func(val interface{}) bool {
	return val == nil || val == xerror.ErrDone || pkg.IsNone(val)
}
