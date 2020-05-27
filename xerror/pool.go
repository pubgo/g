package xerror

import "sync"

var xerrorPool = sync.Pool{New: func() interface{} {
	return &xerror{}
}}
