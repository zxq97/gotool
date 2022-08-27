package redisx

import (
	"github.com/go-redis/redis/v8"
	"github.com/zxq97/gotool/config"
)

type RedisX struct {
	redis redis.Cmdable
}

func NewRedisX(conf *config.RedisConf) *RedisX {
	return &RedisX{
		redis: config.InitRedis(conf),
	}
}
