package memcachex

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/zxq97/gotool/config"
)

type MemcacheX struct {
	*memcache.Client
}

func NewMemcacheX(addr []string) *MemcacheX {
	return &MemcacheX{
		config.InitMC(addr),
	}
}
