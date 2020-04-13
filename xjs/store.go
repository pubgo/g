package xjs

import (
	"github.com/dave/flux"
	"github.com/pubgo/g/pkg"
	"github.com/pubgo/g/xerror"
)

func (t *Vapper) Store(store flux.StoreInterface) {
	xerror.PanicT(pkg.IsNone(store), "store is null")
	t.stores = append(t.stores, store)
}
