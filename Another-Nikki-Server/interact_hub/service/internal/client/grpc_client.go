package client

import (
	"Another-Nikki/interact_hub/service/internal/conf"
	judge "Another-Nikki/judge/service/api"
	"github.com/google/wire"
)

var GrpcProviderSet = wire.NewSet(NewGlobalGrpcClient)

type GlobalGrpcClient struct {
	JudgeClient judge.JudgeClient
}

func NewGlobalGrpcClient(c *conf.Data, timeout *conf.ClientTimeout) *GlobalGrpcClient {
	globalGrpcClient := &GlobalGrpcClient{}
	globalGrpcClient.JudgeClient = judge.NewClientJudge(c.JudgeServiceAddr, timeout.Timeout.AsDuration())
	return globalGrpcClient
}
