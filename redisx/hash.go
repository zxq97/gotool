package redisx

import (
	"context"
	"time"
)

func (rx *RedisX) HIncrByXEX(ctx context.Context, key, field string, incr int64, ttl time.Duration) error {
	ok, err := rx.Expire(ctx, key, ttl).Result()
	if err != nil || !ok {
		return err
	}
	return rx.HIncrByEX(ctx, key, field, incr, ttl)
}

func (rx *RedisX) HIncrByEX(ctx context.Context, key, field string, incr int64, ttl time.Duration) error {
	pipe := rx.Pipeline()
	pipe.HIncrBy(ctx, key, field, incr)
	pipe.Expire(ctx, key, ttl)
	_, err := pipe.Exec(ctx)
	return err
}

func (rx *RedisX) HMSetEX(ctx context.Context, key string, fieldMap map[string]interface{}, ttl time.Duration) error {
	pipe := rx.Pipeline()
	pipe.HMSet(ctx, key, fieldMap)
	pipe.Expire(ctx, key, ttl)
	_, err := pipe.Exec(ctx)
	return err
}
