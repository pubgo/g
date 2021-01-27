package xjs

import (
	"github.com/pubgo/x/pkg"
	"github.com/pubgo/x/xerror"
	"reflect"
)

func (t *Vapper) Config(cfg interface{}) {
	xerror.PanicT(pkg.IsNone(cfg), "please init config")
	t.cfg = reflect.ValueOf(cfg)
}
