package utilx

import "github.com/pubgo/xerror"

func Try(fn func(), catch ...func(err error)) {
	xerror.Assert(fn == nil, "[fn] should not be nil")

	defer xerror.Resp(func(err xerror.XErr) {
		if len(catch) > 0 {
			catch[0](err)
		}
	})

	fn()
}
