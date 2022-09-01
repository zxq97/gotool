package memcachex

import (
	"context"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/zxq97/gotool/concurrent"
)

func (mcx *MemcacheX) SetCtx(ctx context.Context, key string, val []byte, ttl int32) (err error) {
	done := make(chan struct{})
	concurrent.Go(func() {
		defer close(done)
		err = mcx.Set(&memcache.Item{Key: key, Value: val, Expiration: ttl})
	})
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-done:
		return err
	}
}

func (mcx *MemcacheX) GetCtx(ctx context.Context, key string) (val *memcache.Item, err error) {
	done := make(chan struct{})
	concurrent.Go(func() {
		defer close(done)
		val, err = mcx.Get(key)
	})
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-done:
		return val, err
	}
}

func (mcx *MemcacheX) GetMultiCtx(ctx context.Context, keys []string) (val map[string]*memcache.Item, err error) {
	done := make(chan struct{})
	concurrent.Go(func() {
		defer close(done)
		val, err = mcx.GetMulti(keys)
	})
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-done:
		return val, err
	}
}
