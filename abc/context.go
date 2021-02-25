package abc

import (
	"context"
)

type CancelValue struct {
	Cancel context.CancelFunc
	Err    error
}
