package rpc

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/zxq97/gotool/config"
)

func initBreaker(conf *config.HystrixConf) {
	hystrix.ConfigureCommand(conf.Name, hystrix.CommandConfig{
		RequestVolumeThreshold: conf.RequestThreshold,
		MaxConcurrentRequests: conf.MaxRequests,
		Timeout: conf.Timeout,
		SleepWindow: conf.SleepWindow,
		ErrorPercentThreshold: conf.ErrorPercent,
	})
}
