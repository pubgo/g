package xcache_memory

import (
	"github.com/patrickmn/go-cache"
	"github.com/pubgo/x/xcache/xcache_abc"
	"time"
)

const (
	NoExpiration      = cache.NoExpiration
	DefaultExpiration = cache.DefaultExpiration
)

type memoryStore struct {
	c *cache.Cache
}

// NewInMemoryStore returns a InMemoryStore
func New(defaultExpiration time.Duration) xcache_abc.ICache {
	return &memoryStore{c: cache.New(defaultExpiration, time.Minute)}
}

// Get (see CacheStore interface)
func (c *memoryStore) Get(key string) (value interface{}, err error) {
	val, b := c.c.Get(key)
	if !b {
		err = xcache_abc.ErrCacheMiss
	}
	return val, err
}

// Set (see CacheStore interface)
func (c *memoryStore) Set(key string, value interface{}, expires time.Duration) error {
	c.c.Set(key, value, expires)
	return nil
}

func (c *memoryStore) SetDefault(key string, value interface{}) error {
	c.c.SetDefault(key, value)
	return nil
}

// Add (see CacheStore interface)
func (c *memoryStore) Add(key string, value interface{}, expires time.Duration) error {
	return c.c.Add(key, value, expires)
}

func (c *memoryStore) AddDefault(key string, value interface{}) error {
	return c.c.Add(key, value, DefaultExpiration)
}

// Replace (see CacheStore interface)
func (c *memoryStore) Replace(key string, value interface{}, expires time.Duration) error {
	return c.c.Replace(key, value, expires)
}

func (c *memoryStore) ReplaceDefault(key string, value interface{}) error {
	return c.c.Replace(key, value, DefaultExpiration)
}

// Delete (see CacheStore interface)
func (c *memoryStore) Delete(key string) error {
	c.c.Delete(key)
	return nil
}

// Flush (see CacheStore interface)
func (c *memoryStore) Flush() error {
	c.c.Flush()
	return nil
}
