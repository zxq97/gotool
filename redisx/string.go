package redisx

import (
	"context"
	"time"
)

func (rx *RedisX) IncrByXEX(ctx context.Context, key string, incr int64, ttl time.Duration) error {
	ok, err := rx.Expire(ctx, key, ttl).Result()
	if err != nil || !ok {
		return err
	}
	return rx.IncrBy(ctx, key, incr).Err()
}
