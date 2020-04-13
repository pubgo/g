package xcache_abc

import (
	"time"
)

// ICache is the interface of a cache backend
type ICache interface {
	// Set sets an item to the cache, replacing any existing item.
	Set(key string, value interface{}, expire time.Duration) error
	SetDefault(key string, value interface{}) error

	// Add adds an item to the cache only if an item doesn't already exist for the given
	// key, or if the existing item has expired. Returns an error otherwise.
	Add(key string, value interface{}, expire time.Duration) error
	AddDefault(key string, value interface{}) error

	// Delete removes an item from the cache. Does nothing if the key is not in the cache.
	Delete(key string) error

	// Replace sets a new value for the cache key only if it already exists. Returns an
	// error if it does not.
	Replace(key string, data interface{}, expire time.Duration) error
	ReplaceDefault(key string, data interface{}) error

	// Get retrieves an item from the cache. Returns the item or nil, and a bool indicating
	// whether the key was found.
	Get(key string) (val interface{}, err error)
}
