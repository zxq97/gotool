package redisx

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/zxq97/gotool/config"
)

func TestHGetXEX(t *testing.T) {
	rx := NewRedisX(&config.RedisConf{
		Addr: []string{"127.0.0.1:6379"},
	})
	err := rx.ExistEX(context.TODO(), "k", time.Hour)
	if err != nil {
		log.Fatalln(err)
	}
}
