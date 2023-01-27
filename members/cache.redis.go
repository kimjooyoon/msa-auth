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
}
type RdsClient interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Get(ctx context.Context, key string) *redis.StringCmd
}

func NewRedis(r *redis.Client, ctx cache.Context) R {
	return RC{r, ctx}
}

func (r RC) Logout(token string) error {
	err1 := r.rdb.Set(r.ctx, token, "1", jwt.ExpiresTime).Err()
	if err1 != nil {
		return err1
	}

	return nil
}

func (r RC) Valid(token string) error {
	val, err1 := r.rdb.Get(r.ctx, token).Result()
	if val == "" {
		return nil
	}
	if err1 != nil {
		return err1
	}
	if val == "1" {
		return errors.New("token in black list")
	}

	return nil
}
