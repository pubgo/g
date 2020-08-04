package wrapper

import "github.com/pubgo/x/xlog/xlog_core"

func IsDebug() bool {
	return xlog_core.IsDebug
}
