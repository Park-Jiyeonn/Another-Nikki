package data

import (
	"Another-Nikki/code_processing/service/internal/conf"
	judge "Another-Nikki/judge/service/api"
	"context"
	etcdregitry "github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	etcdclient "go.etcd.io/etcd/client/v3"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewRegistry, NewDiscovery, NewGlobalGrpcClient, NewData)

// Data .
type Data struct {
	// TODO wrapped database client
	GlobalGrpcClient *GlobalGrpcClient
}

type GlobalGrpcClient struct {
	JudgeClient judge.JudgeClient
}

func NewDiscovery() registry.Discovery {
	client, err := etcdclient.New(etcdclient.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	})
	if err != nil {
		panic(err)
	}
	return etcdregitry.New(client) // 传入 etcd client，也就是选择 etcd 为服务中心
}

func NewRegistry() registry.Registrar {
	client, err := etcdclient.New(etcdclient.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	})
	if err != nil {
		panic(err)
	}
	return etcdregitry.New(client) // 传入 etcd client，也就是选择 etcd 为服务中心
}

func NewGlobalGrpcClient(c *conf.Data, r registry.Discovery) *GlobalGrpcClient {
	globalGrpcClient := &GlobalGrpcClient{}
	connGRPC, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(c.Registry.Address), // 服务发现
		grpc.WithDiscovery(r),                 // 传入etcd registry
		grpc.WithMiddleware(
			tracing.Client(),
			recovery.Recovery(),
		),
		grpc.WithTimeout(2*time.Second),
		//grpc.WithOptions(grpcx.WithStatsHandler(&tracing.ClientHandler{})),
	)
	if err != nil {
		panic(err)
	}

	globalGrpcClient.JudgeClient = judge.NewJudgeClient(connGRPC)
	return globalGrpcClient
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, globalGrpcClient *GlobalGrpcClient) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		GlobalGrpcClient: globalGrpcClient,
	}, cleanup, nil
}
