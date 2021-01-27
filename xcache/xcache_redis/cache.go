package xcache_redis

import (
	"context"
	"errors"
	"github.com/pubgo/x/xerror"
	"sync/atomic"
	"time"

	"github.com/go-redis/redis/v7"
)

var ErrCacheMiss = errors.New("cache: key is missing")
var errRedisLocalCacheNil = errors.New("cache: both Redis and LocalCache are nil")

type Item struct {
	Ctx context.Context

	Key    string
	Object interface{}

	// Func returns object to be cached.
	Func func() (interface{}, error)

	// Expiration is the cache expiration time.
	// Default expiration is 1 hour.
	Expiration time.Duration
}

func (item *Item) Context() context.Context {
	if item.Ctx == nil {
		return context.Background()
	}
	return item.Ctx
}

func (item *Item) object() (interface{}, error) {
	if item.Object != nil {
		return item.Object, nil
	}
	if item.Func != nil {
		return item.Func()
	}
	return nil, nil
}

func (item *Item) exp() time.Duration {
	if item.Expiration < 0 {
		return 0
	}
	if item.Expiration < time.Second {
		return time.Hour
	}
	return item.Expiration
}

type Codec struct {
	Redis xservice.IRedis

	localCache *xservice.Cache

	Marshal   func(interface{}) ([]byte, error)
	Unmarshal func([]byte, interface{}) error

	group xservice.Group

	hits        uint64
	misses      uint64
	localHits   uint64
	localMisses uint64
}

// UseLocalCache causes Codec to cache items in local LRU cache.
func (cd *Codec) UseLocalCache(maxLen int, expiration time.Duration) {
	cd.localCache = xservice.New(maxLen, expiration)
}

// Set caches the item.
func (cd *Codec) Set(item *Item) error {
	obj, err := item.object()
	if err != nil {
		return err
	}
	_, err = cd.set(item.Context(), item.Key, obj, item.exp())
	return err
}

func (cd *Codec) set(ctx context.Context, key string, obj interface{}, exp time.Duration) (b []byte, err error) {
	defer xerror.RespErr(&err)

	b, err = cd.Marshal(obj)
	xerror.PanicM(err, "cache: Marshal key=%q failed: %s", key, err)

	if cd.localCache != nil {
		cd.localCache.Set(key, b)
	}

	if cd.Redis == nil {
		if cd.localCache == nil {
			return nil, errRedisLocalCacheNil
		}
		return b, nil
	}

	err = cd.Redis.Set(key, b, exp).Err()
	xerror.PanicM(err, "cache: Set key=%q failed: %s", key, err)

	return b, err
}

// Exists reports whether object for the given key exists.
func (cd *Codec) Exists(key string) bool {
	return cd.Get(key, nil) == nil
}

func (cd *Codec) ExistsContext(ctx context.Context, key string) bool {
	return cd.GetContext(ctx, key, nil) == nil
}

// Get gets the object for the given key.
func (cd *Codec) Get(key string, object interface{}) error {
	return cd.get(context.Background(), key, object, false)
}

func (cd *Codec) GetContext(ctx context.Context, key string, object interface{}) error {
	return cd.get(ctx, key, object, false)
}

func (cd *Codec) get(ctx context.Context, key string, object interface{}, onlyLocalCache bool, ) error {
	b, err := cd.getBytes(key, onlyLocalCache)
	if err != nil {
		return err
	}

	if object == nil || len(b) == 0 {
		return nil
	}

	err = cd.Unmarshal(b, object)
	xerror.PanicM(err, "cache: key=%q Unmarshal(%T) failed: %s", key, object, err)

	return nil
}

func (cd *Codec) getBytes(key string, onlyLocalCache bool) ([]byte, error) {
	if cd.localCache != nil {
		b, ok := cd.localCache.Get(key)
		if ok {
			atomic.AddUint64(&cd.localHits, 1)
			return b, nil
		}
		atomic.AddUint64(&cd.localMisses, 1)
	}

	if onlyLocalCache {
		return nil, ErrCacheMiss
	}
	if cd.Redis == nil {
		if cd.localCache == nil {
			return nil, errRedisLocalCacheNil
		}
		return nil, ErrCacheMiss
	}

	b, err := cd.Redis.Get(key).Bytes()
	if err != nil {
		atomic.AddUint64(&cd.misses, 1)
		if err == redis.Nil {
			return nil, ErrCacheMiss
		}

		xerror.PanicM(err, "cache: Get key=%q failed: %s", key, err)
		return nil, err
	}
	atomic.AddUint64(&cd.hits, 1)

	if cd.localCache != nil {
		cd.localCache.Set(key, b)
	}
	return b, nil
}

// Once gets the item.Object for the given item.Key from the cache or
// executes, caches, and returns the results of the given item.Func,
// making sure that only one execution is in-flight for a given item.Key
// at a time. If a duplicate comes in, the duplicate caller waits for the
// original to complete and receives the same results.
func (cd *Codec) Once(item *Item) error {
	b, cached, err := cd.getSetItemBytesOnce(item)
	if err != nil {
		return err
	}

	if item.Object == nil || len(b) == 0 {
		return nil
	}

	err = cd.Unmarshal(b, item.Object)
	if err != nil {
		if cached {
			_ = cd.Delete(item.Key)
			return cd.Once(item)
		}
		xerror.PanicM(err, "cache: key=%q Unmarshal(%T) failed: %s", item.Key, item.Object, err)
		return err
	}

	return nil
}

func (cd *Codec) getSetItemBytesOnce(item *Item) (b []byte, cached bool, err error) {
	if cd.localCache != nil {
		b, err := cd.getItemBytesFast(item)
		if err == nil {
			return b, true, nil
		}
	}

	obj, err := cd.group.Do(item.Key, func() (interface{}, error) {
		b, err := cd.getItemBytes(item)
		if err == nil {
			cached = true
			return b, nil
		}

		obj, err := item.Func()
		if err != nil {
			return nil, err
		}

		b, err = cd.set(item.Context(), item.Key, obj, item.exp())
		if b != nil {
			// Ignore error if we have the result.
			return b, nil
		}
		return nil, err
	})
	if err != nil {
		return nil, false, err
	}
	return obj.([]byte), cached, nil
}

func (cd *Codec) getItemBytes(item *Item) ([]byte, error) {
	return cd.getBytes(item.Key, false)
}

func (cd *Codec) getItemBytesFast(item *Item) ([]byte, error) {
	return cd.getBytes(item.Key, true)
}

func (cd *Codec) Delete(key string) error {
	return cd.DeleteContext(context.Background(), key)
}

func (cd *Codec) DeleteContext(ctx context.Context, key string) error {
	if cd.localCache != nil {
		cd.localCache.Delete(key)
	}

	if cd.Redis == nil {
		if cd.localCache == nil {
			return errRedisLocalCacheNil
		}
		return nil
	}

	deleted, err := cd.Redis.Del(key).Result()
	xerror.PanicM(err, "cache: Del key=%q failed: %s", key, err)

	if deleted == 0 {
		return ErrCacheMiss
	}
	return nil
}

//------------------------------------------------------------------------------

type Stats struct {
	Hits        uint64
	Misses      uint64
	LocalHits   uint64
	LocalMisses uint64
}

// Stats returns cache statistics.
func (cd *Codec) Stats() *Stats {
	stats := Stats{
		Hits:   atomic.LoadUint64(&cd.hits),
		Misses: atomic.LoadUint64(&cd.misses),
	}
	if cd.localCache != nil {
		stats.LocalHits = atomic.LoadUint64(&cd.localHits)
		stats.LocalMisses = atomic.LoadUint64(&cd.localMisses)
	}
	return &stats
}
