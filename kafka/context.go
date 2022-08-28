package kafka

import (
	"context"

	"github.com/zxq97/gotool/constant"
	"github.com/zxq97/gotool/generate"
)

func consumerContext(traceid string) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultConsumerTimeout)
	if traceid == "" {
		traceid = generate.UUIDStr()
	}
	ctx = context.WithValue(ctx, constant.TraceIDKey, traceid)
	return ctx, cancel
}
