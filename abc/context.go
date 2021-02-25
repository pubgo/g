package abc

import (
	"context"
)

type Cancel struct {
	Cancel context.CancelFunc
	Err    error
}
