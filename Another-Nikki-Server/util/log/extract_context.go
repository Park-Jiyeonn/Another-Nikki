package log

import (
	"context"
	"google.golang.org/grpc/metadata"
)

func ExtractTraceIDFromContext(ctx context.Context) (string, bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", false
	}

	values := md.Get("trace-id")
	if len(values) == 0 {
		return "", false
	}

	return values[0], true
}

func AddTraceIDToContext(ctx context.Context, traceID string) context.Context {
	md := metadata.Pairs("trace-id", traceID)
	return metadata.NewOutgoingContext(ctx, md)
}
