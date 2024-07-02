package api

import (
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"golang.org/x/net/context"
	"time"
)

const JudgeAppID = "Another-Nikki.Judge-Service"

func NewClientJudge(r registry.Discovery, timeout time.Duration) JudgeClient {
	connGRPC, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery://default/"+JudgeAppID), // 服务发现
		grpc.WithDiscovery(r), // 传入etcd registry
		grpc.WithMiddleware(
			tracing.Client(),
			recovery.Recovery(),
		),
		grpc.WithTimeout(timeout), // 虽然 DialInsecure 中默认超时时间是 2s，但是这里传入的会对其进行修改
		//grpc.WithOptions(grpcx.WithStatsHandler(&tracing.ClientHandler{})),
	)
	if err != nil {
		panic(err)
	}
	return NewJudgeClient(connGRPC)
}
