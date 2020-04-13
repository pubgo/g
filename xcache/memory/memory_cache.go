package memory

import (
	"errors"
	"github.com/patrickmn/go-cache"
	"reflect"
	"time"
)

const (
	NoExpiration      = cache.NoExpiration
	DefaultExpiration = cache.DefaultExpiration
)

func New(defaultExpiration, cleanupInterval time.Duration) *cache.Cache {
	return cache.New(defaultExpiration, cleanupInterval)
}

//InMemoryStore represents the cache with memory persistence
type InMemoryStore struct {
	cache.Cache
}

// NewInMemoryStore returns a InMemoryStore
func NewInMemoryStore(defaultExpiration time.Duration) *InMemoryStore {
	return &InMemoryStore{*cache.New(defaultExpiration, time.Minute)}
}

// Get (see CacheStore interface)
func (c *InMemoryStore) Get(key string, value interface{}) error {
	val, found := c.Cache.Get(key)
	if !found {
		return xservice.ErrCacheMiss
	}

	v := reflect.ValueOf(value)
	if v.Type().Kind() == reflect.Ptr && v.Elem().CanSet() {
		v.Elem().Set(reflect.ValueOf(val))
		return nil
	}
	return xservice.ErrNotStored
}

// Set (see CacheStore interface)
func (c *InMemoryStore) Set(key string, value interface{}, expires time.Duration) error {
	// NOTE: go-cache understands the values of DEFAULT and FOREVER
	c.Cache.Set(key, value, expires)
	return nil
}

// Add (see CacheStore interface)
func (c *InMemoryStore) Add(key string, value interface{}, expires time.Duration) error {
	return c.Cache.Add(key, value, expires)
}

// Replace (see CacheStore interface)
func (c *InMemoryStore) Replace(key string, value interface{}, expires time.Duration) error {
	if err := c.Cache.Replace(key, value, expires); err != nil {
		return xservice.ErrNotStored
	}
	return nil
}

// Delete (see CacheStore interface)
func (c *InMemoryStore) Delete(key string) error {
	c.Cache.Delete(key)
	return nil
}

// Increment (see CacheStore interface)
func (c *InMemoryStore) Increment(key string, n int64) error {
	if c.Cache.Increment(key, n) == xservice.ErrCacheMiss {
		return xservice.ErrCacheMiss
	}
	return nil
}

// Decrement (see CacheStore interface)
func (c *InMemoryStore) Decrement(key string, n int64) error {
	if c.Cache.Decrement(key, n) == xservice.ErrCacheMiss {
		return xservice.ErrCacheMiss
	}
	return nil
}

// Flush (see CacheStore interface)
func (c *InMemoryStore) Flush() error {
	c.Cache.Flush()
	return nil
}

func (c *InMemoryStore) Expire(key string, expires time.Duration) (bool, error) {
	return false, errors.New("not implemented")
}
