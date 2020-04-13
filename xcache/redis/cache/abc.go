package cache

import (
	"github.com/go-redis/redis/v7"
	"time"
)

type IRedis interface {
	Get(key string) *redis.StringCmd
	Del(keys ...string) *redis.IntCmd
	Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd
}
