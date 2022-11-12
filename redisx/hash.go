package redisx

import (
	"context"
	"time"
)

func (rx *RedisX) HIncrByXEX(ctx context.Context, key, field string, incr int64, ttl time.Duration) error {
	if err := rx.ExistEX(ctx, key, ttl); err != nil {
		return err
	}
	return rx.HIncrBy(ctx, key, field, incr).Err()
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

func (rx *RedisX) HMSetXEX(ctx context.Context, key string, fieldMap map[string]interface{}, ttl time.Duration) error {
	if err := rx.ExistEX(ctx, key, ttl); err != nil {
		return err
	}
	return rx.HMSet(ctx, key, fieldMap).Err()
}

func (rx *RedisX) HMGetXEX(ctx context.Context, key string, ttl time.Duration, field ...string) ([]interface{}, error) {
	if err := rx.ExistEX(ctx, key, ttl); err != nil {
		return nil, err
	}
	return rx.HMGet(ctx, key, field...).Result()
}

func (rx *RedisX) HGetXEX(ctx context.Context, key, field string, ttl time.Duration) (string, error) {
	if err := rx.ExistEX(ctx, key, ttl); err != nil {
		return "", err
	}
	return rx.HGet(ctx, key, field).Result()
}
