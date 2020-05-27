package xerror

import "sync"

var xerrorPool = sync.Pool{New: func() interface{} {
	return &xerror{}
}}

func getXerror() *xerror {
	return xerrorPool.Get().(*xerror)
}

func putXerror(err *xerror) {
	xerrorPool.Put(err)
}
