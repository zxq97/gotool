package redisx

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/zxq97/gotool/cast"
	"github.com/zxq97/gotool/concurrent"
	"github.com/zxq97/gotool/constant"
	"github.com/zxq97/gotool/sequence"
)

var (
	zRevRangeByMemberScript = redis.NewScript(`
        local rank = redis.call("ZRevRank", KEYS[1], ARGV[1])
        if not rank then
            rank = 0
		else
			rank = rank + 1
        end
        return redis.call("ZRevRange", KEYS[1], rank, rank + ARGV[2])
    `)

	zRevRangeByMemberWithScoresScript = redis.NewScript(`
		local rank = redis.call("ZRevRank", KEYS[1], ARGV[1])
		if not rank then
			rank = 0
		else
			rank = rank + 1
		end
		return redis.call("ZRevRange", KEYS[1], rank, rank + ARGV[2], 'withscores')
	`)
)

func (rx *RedisX) ZAddXEX(ctx context.Context, key string, zs []*redis.Z, ttl time.Duration) error {
	if err := rx.ExistEX(ctx, key, ttl); err != nil {
		return err
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

func (rx *RedisX) ZRevRangeByMember(ctx context.Context, key, member string, offset int64) ([]int64, error) {
	res, err := zRevRangeByMemberScript.Run(ctx, rx, []string{key}, member, offset).Result()
	if err != nil {
		return nil, err
	}
	val, ok := res.([]interface{})
	if !ok || len(val) == 0 {
		return nil, redis.Nil
	}
	ids := make([]int64, 0, offset)
	for _, v := range val {
		id := v.(string)
		ids = append(ids, cast.ParseInt(id, 0))
	}
	return ids, nil
}

func (rx *RedisX) ZRevRangeByMemberWithScores(ctx context.Context, key, member string, offset int64) ([]*redis.Z, error) {
	res, err := zRevRangeByMemberWithScoresScript.Run(ctx, rx, []string{key}, member, offset).Result()
	if err != nil {
		return nil, err
	}
	val, ok := res.([]interface{})
	if !ok || len(val) == 0 || len(val)&1 != 0 {
		return nil, redis.Nil
	}
	zs := make([]*redis.Z, 0, len(val)>>1)
	for i := 0; i < len(val); i += 2 {
		id := val[i+1].(string)
		zs = append(zs, &redis.Z{Member: val[i], Score: float64(cast.ParseInt(id, 0))})
	}
	return zs, nil
}
