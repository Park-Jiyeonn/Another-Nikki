package service

import (
	"Another-Nikki/service/judge/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type globalGrpcClient struct {
	JudgeClient api.JudgeClient
}

type Service struct {
	GlobalGrpc *globalGrpcClient
}

func New() (s *Service) {
	s = &Service{}
	glbgrpc := &globalGrpcClient{}
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	glbgrpc.JudgeClient = api.NewJudgeClient(conn)
	s.GlobalGrpc = glbgrpc
	return
}
