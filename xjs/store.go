package xjs

import (
	"github.com/dave/flux"
	"github.com/pubgo/x/pkg"
	"github.com/pubgo/xerror"
)

func (t *Vapper) Store(store flux.StoreInterface) {
	xerror.PanicT(pkg.IsNone(store), "store is null")
	t.stores = append(t.stores, store)
}
