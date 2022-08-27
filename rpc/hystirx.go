package rpc

import "github.com/afex/hystrix-go/hystrix"

func initBreaker(commandName string) {
	hystrix.ConfigureCommand(commandName, hystrix.CommandConfig{
		RequestVolumeThreshold: 5,
		MaxConcurrentRequests: 100,
		Timeout: 3000,
		SleepWindow: 10000,
		ErrorPercentThreshold: 20,
	})
}
