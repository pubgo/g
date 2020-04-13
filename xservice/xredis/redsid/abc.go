package redsid

import (
	"github.com/go-redis/redis/v7"
	"time"
)

type IRedis interface {
	SetNX(key string, value interface{}, expiration time.Duration) *redis.BoolCmd
}
