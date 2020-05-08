package xerror

import "errors"

var (
	// ErrDone done
	ErrDone = errors.New("DONE")
	// Err base error
	Err = NewXErr("XErr")
	// ErrUnknownType error
	ErrUnknownType = Err.New("UnknownType")
)
