package api

import (
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"golang.org/x/net/context"
	"time"
)

const JudgeAppID = "Another-Nikki.Judge-Service"

func NewClientJudge(addr string, timeout time.Duration) JudgeClient {
	connGRPC, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(addr),
		grpc.WithMiddleware(
			tracing.Client(),
			recovery.Recovery(),
		),
		grpc.WithTimeout(timeout),
	)
	if err != nil {
		panic(err)
	}
	return NewJudgeClient(connGRPC)
}
