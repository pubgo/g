package try

import (
	"github.com/pubgo/x/stack"
	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog"
	"go.uber.org/zap"
)

func Logs(logs xlog.Xlog, fn func(), fields ...zap.Field) {
	xerror.Assert(fn == nil, "[fn] should not be nil")
	xerror.Assert(logs == nil, "[logs] should not be nil")

	defer xerror.Resp(func(err xerror.XErr) {
		var params = make([]interface{}, 0, len(fields)+2)
		for i := range fields {
			params = append(params, fields[i])
		}
		logs.Error(append(params, zap.Any("err", err), zap.String("stack", stack.Func(fn)))...)
	})

	fn()
}

func Catch(fn func(), catch ...func(err error)) {
	xerror.Assert(fn == nil, "[fn] should not be nil")

	if len(catch) > 0 && catch[0] != nil {
		defer xerror.Resp(func(err xerror.XErr) { catch[0](err) })
	}

	fn()
}

func With(err *error, fn func()) {
	xerror.Assert(fn == nil, "[fn] should not be nil")

	defer xerror.RespErr(err)
	fn()

	return
}

func Try(fn func()) (err error) {
	xerror.Assert(fn == nil, "[fn] should not be nil")

	defer xerror.RespErr(&err)

	fn()
	return
}
