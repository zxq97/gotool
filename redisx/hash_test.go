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

func TestRedisX_HMSetEX(t *testing.T) {
	rx := NewRedisX(&config.RedisConf{
		Addr: []string{"127.0.0.1:6379"},
	})
	fieldMap := map[string]interface{}{
		"1": "1",
		"2": "2",
	}
	err := rx.HMSetEX(context.TODO(), "h", fieldMap, time.Hour)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestRedisX_HMSetXEX(t *testing.T) {
	rx := NewRedisX(&config.RedisConf{
		Addr: []string{"127.0.0.1:6379"},
	})
	fieldMap := map[string]interface{}{
		"1": "1",
		"2": "2",
	}
	err := rx.HMSetXEX(context.TODO(), "h", fieldMap, time.Hour)
	if err != nil {
		log.Fatalln(err)
	}
}

func TestRedisX_HMGetXEX(t *testing.T) {
	rx := NewRedisX(&config.RedisConf{
		Addr: []string{"127.0.0.1:6379"},
	})
	val, err := rx.HMGetXEX(context.TODO(), "h", time.Hour, "1", "2", "3")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(val)
}

func TestRedisX_HGetXEX(t *testing.T) {
	rx := NewRedisX(&config.RedisConf{
		Addr: []string{"127.0.0.1:6379"},
	})
	val, err := rx.HGetXEX(context.TODO(), "h", "1", time.Hour)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(val)
}
