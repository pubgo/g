package future

import (
	"github.com/pubgo/x/abc"
)

var _ abc.Future = (*future)(nil)

type future struct{ p *promise }

func (f *future) Yield(data interface{})                      { f.p.yield(data) }
func (f *future) YieldFn(val abc.FutureValue, fn interface{}) { f.p.await(val, fn) }
