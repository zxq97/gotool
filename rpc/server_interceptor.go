package rpc

import (
	"context"
	"log"
	"runtime/debug"
	"time"

	"github.com/zxq97/gotool/constant"
	"google.golang.org/grpc"
)

func recovery(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	traceid := getIncomingReqID(ctx)
	if traceid != "" {
		ctx = context.WithValue(ctx, constant.TraceIDKey, traceid)
	}
	now := time.Now()
	defer func() {
		log.Println(time.Since(now))
		if err := recover(); err != nil {
			log.Println(info.FullMethod, err, string(debug.Stack()))
		}
	}()
	return handler(ctx, req)
}
