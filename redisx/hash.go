package redisx

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

func (rx *RedisX) HIncrByXEX(ctx context.Context, key, field string, incr int64, ttl time.Duration) error {
	ok, err := rx.Expire(ctx, key, ttl).Result()
	if err != nil {
		return err
	} else if !ok {
		return redis.Nil
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
	ok, err := rx.Expire(ctx, key, ttl).Result()
	if err != nil {
		return err
	} else if !ok {
		return redis.Nil
	}
	return rx.HMSet(ctx, key, fieldMap).Err()
}

func (rx *RedisX) HMGetXEX(ctx context.Context, key string, ttl time.Duration, field ...string) ([]interface{}, error) {
	ok, err := rx.Expire(ctx, key, ttl).Result()
	if err != nil {
		return nil, err
	} else if !ok {
		return nil, redis.Nil
	}
	return rx.HMGet(ctx, key, field...).Result()
}
