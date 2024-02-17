package data

import (
	codeService "Another-Nikki/interact_hub/service/api"
	"Another-Nikki/judge/job/internal/conf"
	judge "Another-Nikki/judge/service/api"
	etcdregitry "github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/registry"
	etcdclient "go.etcd.io/etcd/client/v3"
)

type GlobalGrpcClient struct {
	JudgeClient          judge.JudgeClient
	CodeProcessingClient codeService.CodeProcessingClient
}

func NewGlobalGrpcClient(c *conf.Data, r registry.Discovery) *GlobalGrpcClient {
	globalGrpcClient := &GlobalGrpcClient{}
	globalGrpcClient.JudgeClient = judge.NewClientJudge(r)
	globalGrpcClient.CodeProcessingClient = codeService.NewClientCodeProcessing(r)
	return globalGrpcClient
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
