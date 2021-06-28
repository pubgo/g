package try

import (
	"github.com/pubgo/x/stack"
	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog"
	"go.uber.org/zap"
)

func Logs(fn func(), logs ...xlog.Xlog) {
	xerror.Assert(fn == nil, "[fn] should not be nil")

	if len(logs) > 0 && logs[0] != nil {
		defer xerror.Resp(func(err xerror.XErr) { logs[0].Error(stack.Func(fn), zap.Any("err", err)) })
	}

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
