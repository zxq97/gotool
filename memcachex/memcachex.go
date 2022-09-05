package memcachex

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/zxq97/gotool/config"
)

type MemcacheX struct {
	*memcache.Client
}

func NewMemcacheX(conf *config.MCConf) *MemcacheX {
	return &MemcacheX{
		conf.InitMC(),
	}
}
