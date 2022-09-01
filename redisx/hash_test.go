package redisx

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/zxq97/gotool/config"
)

func TestRedisX_HIncrByEX(t *testing.T) {
	rx := NewRedisX(&config.RedisConf{
		Addr: []string{"127.0.0.1:6379"},
	})
	err := rx.HIncrByEX(context.TODO(), "h", "1", 10, time.Hour)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestRedisX_HIncrByXEX(t *testing.T) {
	rx := NewRedisX(&config.RedisConf{
		Addr: []string{"127.0.0.1:6379"},
	})
	err := rx.HIncrByXEX(context.TODO(), "h", "1", 10, time.Hour)
	if err != nil {
		log.Fatalln(err)
	}
}
