package service

import (
	"Another-Nikki/service/judge-service/api"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	api.UnimplementedJudgeServer
}

func New() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	s := &server{}
	api.RegisterJudgeServer(grpcServer, s)
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
