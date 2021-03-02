package abc

import (
	"context"
)

type Cancel struct {
	Cancel context.CancelFunc
	Err    error
}

func (t Cancel) Error() string {
	if t.Err == nil {
		return ""
	}

	return t.Err.Error()
}
