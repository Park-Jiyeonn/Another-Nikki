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
