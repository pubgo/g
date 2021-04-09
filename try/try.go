package try

import (
	"github.com/pubgo/x/fx"
	"github.com/pubgo/x/stack"
	"github.com/pubgo/xerror"
)

func Catch(fn func(), catch ...func(err error)) {
	xerror.Assert(fn == nil, "[fn] should not be nil")

	defer xerror.Resp(func(err xerror.XErr) {
		if len(catch) == 0 {
			return
		}
		catch[0](xerror.Wrap(err, stack.Func(fn)))
	})

	fn()
}

func With(gErr *error, fn interface{}, args ...interface{}) {
	xerror.Assert(fn == nil, "[fn] should not be nil")

	defer xerror.Resp(func(err xerror.XErr) { *gErr = xerror.Wrap(err, stack.Func(fn)) })
	_ = fx.WrapValue(fn, args...)

	return
}

func Do(fn interface{}, args ...interface{}) (gErr error) {
	xerror.Assert(fn == nil, "[fn] should not be nil")

	defer xerror.Resp(func(err xerror.XErr) { gErr = xerror.Wrap(err, stack.Func(fn)) })
	_ = fx.WrapValue(fn, args...)

	return
}
