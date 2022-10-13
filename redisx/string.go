package redisx

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

func (rx *RedisX) IncrByXEX(ctx context.Context, key string, incr int64, ttl time.Duration) error {
	ok, err := rx.Expire(ctx, key, ttl).Result()
	if err != nil {
		return err
	} else if !ok {
		return redis.Nil
	}
	return rx.IncrBy(ctx, key, incr).Err()
}
