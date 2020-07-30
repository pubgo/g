package xtry

import "errors"

var (
	ErrNotFuncType    = errors.New("func type not match error")
	ErrParamsNotMatch = errors.New("params not match error")
)
