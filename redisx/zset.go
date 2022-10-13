package redisx

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/zxq97/gotool/concurrent"
	"github.com/zxq97/gotool/constant"
	"github.com/zxq97/gotool/sequence"
)

func (rx *RedisX) ZAddXEX(ctx context.Context, key string, zs []*redis.Z, ttl time.Duration) error {
	ok, err := rx.Expire(ctx, key, ttl).Result()
	if err != nil {
		return err
	} else if !ok {
		return redis.Nil
	}
	return rx.ZAddEX(ctx, key, zs, ttl)
}

func (rx *RedisX) ZAddEX(ctx context.Context, key string, zs []*redis.Z, ttl time.Duration) error {
	zss := sequence.Chunks[*redis.Z](zs, constant.DefaultBatchSize)
	eg := concurrent.NewErrGroup(ctx)
	for _, zz := range zss {
		z := zz
		eg.Go(func() error {
			pipe := rx.Pipeline()
			pipe.ZAdd(ctx, key, z...)
			pipe.Expire(ctx, key, ttl)
			_, err := pipe.Exec(ctx)
			return err
		})
	}
	return eg.Wait()
}
