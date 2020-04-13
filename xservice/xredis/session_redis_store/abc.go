package session_redis_store

import (
	"github.com/gin-contrib/sessions"
	"github.com/go-redis/redis/v7"
	"time"
)

type IStore interface {
	sessions.Store
}

type _IRedis interface {
	Close() error
	Get(key string) *redis.StringCmd
	Del(keys ...string) *redis.IntCmd
	Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	SetNX(key string, value interface{}, expiration time.Duration) *redis.BoolCmd
}
