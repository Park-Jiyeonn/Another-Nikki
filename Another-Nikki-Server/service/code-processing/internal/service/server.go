package service

import (
	"Another-Nikki/service/code-processing/api"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	api.UnimplementedCodeProcessingServer
}

func New() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9001))
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	s := &server{}
	api.RegisterCodeProcessingServer(grpcServer, s)
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
