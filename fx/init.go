package fx

import (
	"context"
	"github.com/pubgo/xlog"
)

var logs = xlog.GetLogger("fx")

type Ctx struct {
	context.Context
}

func (Ctx) Break() { panic(errBreak) }
