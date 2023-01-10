package cache

import (
	context2 "context"
	"github.com/go-redis/redis/v9"
	"os"
)

var rds *redis.Client

type Context context2.Context

var ctx Context

func NewRedisContext() Context {
	if ctx != nil {
		return ctx
	}
	ctx = context2.Background()
	return ctx
}

type DSN string

func RedisNewDSN() DSN {
	return DSN(os.Getenv("REDIS_DSN"))
}

func RedisConnection() *redis.Client {
	if rds != nil {
		return rds
	}

	rds = redis.NewClient(&redis.Options{
		Addr:     string(RedisNewDSN()),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rds
}
