package api

import (
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"golang.org/x/net/context"
)

const CodeProcessingAppID = "Another-Nikki.InteractHub-Service"

func NewClientProblem(r registry.Discovery) ProblemClient {
	connGRPC, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery://default/"+CodeProcessingAppID), // 服务发现
		grpc.WithDiscovery(r), // 传入etcd registry
		grpc.WithMiddleware(
			tracing.Client(),
			recovery.Recovery(),
		),
		//grpc.WithTimeout(2*time.Second),
		//grpc.WithOptions(grpcx.WithStatsHandler(&tracing.ClientHandler{})),
	)
	if err != nil {
		panic(err)
	}
	return NewProblemClient(connGRPC)
}
