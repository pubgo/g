package xtry

import "github.com/pubgo/g/xerror"

var (
	Err               = xerror.NewXErr("XTry")
	ErrNotFuncType    = Err.New("func type not match error")
	ErrParamsNotMatch = Err.New("params not match error")
)
