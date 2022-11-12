package redisx

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

func (rx *RedisX) ExistEX(ctx context.Context, key string, ttl time.Duration) error {
	ok, err := rx.Expire(ctx, key, ttl).Result()
	if err != nil {
		return err
	} else if !ok {
		return redis.Nil
	}
	return nil
}
