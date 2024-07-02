package api

import (
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"golang.org/x/net/context"
	"time"
)

const InteractHubAppID = "Another-Nikki.Interact-Service"

func NewClientCodeProcessing(r registry.Discovery, timeout time.Duration) CodeProcessingClient {
	connGRPC, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery://default/"+InteractHubAppID), // 服务发现
		grpc.WithDiscovery(r), // 传入etcd registry
		grpc.WithMiddleware(
			tracing.Client(),
			recovery.Recovery(),
		),
		grpc.WithTimeout(timeout),
		//grpc.WithOptions(grpcx.WithStatsHandler(&tracing.ClientHandler{})),
	)
	if err != nil {
		panic(err)
	}
	return NewCodeProcessingClient(connGRPC)
}
