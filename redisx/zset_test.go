package redisx

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/zxq97/gotool/config"
)

func TestNewRedisX(t *testing.T) {
	rx := NewRedisX(&config.RedisConf{
		Addr: []string{"127.0.0.1:6379"},
	})
	zs := make([]*redis.Z, 0, 200)
	for i := 0; i < 200; i++ {
		zs = append(zs, &redis.Z{
			Member: i,
			Score: float64(i),
		})
	}
	err := rx.ZAddXEX(context.TODO(), "k", zs, time.Hour)
	if err != nil {
		log.Fatalln(err)
	}
}
