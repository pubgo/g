package xtask

import (
	"context"
	"github.com/pubgo/xerror"
	"github.com/pubgo/x/xtry"
	"sync"
	"sync/atomic"
)

type Group struct {
	cancel func()

	wg sync.WaitGroup

	err error

	done uint32
	m    *sync.Mutex
}

func NewGroupTask() *Group {
	return &Group{}
}

func WithContext(ctx context.Context) (*Group, context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	return &Group{cancel: cancel}, ctx
}

func (g *Group) Wait() error {
	g.wg.Wait()
	if g.cancel != nil {
		g.cancel()
	}
	return g.err
}

func (g *Group) Go(fn func()) {
	if atomic.LoadUint32(&g.done) == 1 {
		return
	}

	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		xerror.ErrHandle(xtry.Try(fn), func(err xerror.IErr) {
			// 遇到错误, 停止执行
			if atomic.LoadUint32(&g.done) == 1 {
				return
			}

			g.m.Lock()
			defer g.m.Unlock()

			if g.done == 0 {
				defer atomic.StoreUint32(&g.done, 1)
				g.err = err
				if g.cancel != nil {
					g.cancel()
				}

				xerror.P(err, "group error, method:group")
			}
		})
	}()
}
