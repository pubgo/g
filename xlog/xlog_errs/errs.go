package xlog_errs

import "github.com/pubgo/xerror"

var (
	Err              = xerror.New("xlog err")
	ErrParamsInValid = Err.New("params is invalid")
)
