package members

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v9"
	"msa-auth/cache"
	"msa-auth/util/jwt"
	"time"
)

type R interface {
	Logout(token string) error
	Valid(token string) error
}

type RC struct {
	rdb RdsClient
	ctx cache.Context
	cacheValid
}

type cacheValid interface {
	isOne(string) bool
	isError(error) bool
	err(e error) error
}

type cacheValidImpl struct{}

type RdsClient interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Get(ctx context.Context, key string) *redis.StringCmd
}

func NewRedis(r *redis.Client, ctx cache.Context) R {
	return RC{r, ctx, cacheValidImpl{}}
}

func (r RC) Logout(token string) error {
	return r.rdb.Set(r.ctx, token, "1", jwt.ExpiresTime).Err()
}

func (cacheValidImpl) isOne(s string) bool  { return s == "1" }
func (cacheValidImpl) isError(e error) bool { return e != nil }
func (cacheValidImpl) err(e error) error    { return e }

func (r RC) Valid(token string) error {
	val, err1 := r.rdb.Get(r.ctx, token).Result()
	if r.isError(err1) {
		return r.err(err1)
	}

	if r.isOne(val) {
		return errors.New("token in black list")
	}

	return nil
}
