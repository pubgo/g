package xjs

import (
	"github.com/pubgo/g/pkg"
	"github.com/pubgo/g/xerror"
	"reflect"
)

func (t *Vapper) Config(cfg interface{}) {
	xerror.PanicT(pkg.IsNone(cfg), "please init config")
	t.cfg = reflect.ValueOf(cfg)
}
