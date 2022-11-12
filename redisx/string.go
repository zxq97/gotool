package redisx

import (
	"context"
	"time"
)

func (rx *RedisX) IncrByXEX(ctx context.Context, key string, incr int64, ttl time.Duration) error {
	if err := rx.ExistEX(ctx, key, ttl); err != nil {
		return err
	}
	return rx.IncrBy(ctx, key, incr).Err()
}
