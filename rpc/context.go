package rpc

import (
	"context"
	"time"

	"github.com/zxq97/gotool/constant"
	"github.com/zxq97/gotool/generate"
	"google.golang.org/grpc/metadata"
)

func defaultTimeout(ctx context.Context, d time.Duration) (context.Context, context.CancelFunc) {
	var cancel context.CancelFunc
	if _, ok := ctx.Deadline(); !ok {
		ctx, cancel = context.WithTimeout(ctx, d)
	}
	return ctx, cancel
}

func withOutgoing(ctx context.Context) context.Context {
	rawid := ctx.Value(constant.TraceIDKey)
	if rawid != nil {
		if traceid, ok := rawid.(string); ok {
			return metadata.AppendToOutgoingContext(ctx, constant.TraceIDKey, traceid)
		}
	}
	return metadata.AppendToOutgoingContext(ctx, constant.TraceIDKey, generate.UUIDStr())
}

func getIncomingReqID(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		rs := md.Get(constant.TraceIDKey)
		if len(rs) > 0 {
			return rs[0]
		}
	}
	return ""
}
