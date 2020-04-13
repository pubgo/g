package session_redis_store

import "time"

// Amount of time for cookies/redis keys to expire.
var sessionExpire = time.Hour * 24 * 30
