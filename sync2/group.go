package sync2

import (
	"context"

	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog"
)

type Group = group
type group struct {
	ctx    context.Context
	cancel context.CancelFunc
	wg     WaitGroup
}

func Parallel(fns ...func(ctx context.Context)) {
	var g = NewGroup()
	defer g.Wait()
	for i := 0; i < len(fns); i++ {
		g.Go(fns[i])
	}
}

// NewGroup
// 创建一个group对象, 可以带上默认的Context
func NewGroup(c ...uint16) *group {
	ctx, cancel := context.WithCancel(context.Background())
	g := &group{ctx: ctx, cancel: cancel, wg: NewWaitGroup(true, c...)}
	return g
}

// Cancel
// 停止正在运行的函数
func (g *group) Cancel() { g.cancel() }

// Count
// 当前的goroutine数量
func (g *group) Count() uint16 { return g.wg.Count() }

// Wait
// 等待正在运行的函数
func (g *group) Wait() { g.wg.Wait(); g.cancel() }

// Go
// 运行一个goroutine
func (g *group) Go(fn func(ctx context.Context)) {
	xerror.Assert(fn == nil, "[fn] should not be nil")

	g.wg.Inc()
	go func() {
		defer g.wg.Done()
		defer xerror.Resp(func(err xerror.XErr) { xlog.Error("group.Go error", xlog.Any("err", err)) })
		fn(g.ctx)
	}()
}
