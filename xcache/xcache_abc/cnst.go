package xcache_abc

import (
	"errors"
	"time"
)

const (
	DEFAULT = time.Duration(0)
	FOREVER = time.Duration(-1)
)

var (
	KeyPrefix     = "golug.xcache"
	ErrCacheMiss  = errors.New("cache: key not found")
	ErrNotStored  = errors.New("cache: not stored")
	ErrNotSupport = errors.New("cache: not support")
)
