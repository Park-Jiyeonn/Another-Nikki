package service

import (
	"context"

	pb "Another-Nikki/code_processing/service/api"
)

type CodeProcessingService struct {
	pb.UnimplementedCodeProcessingServer
}

func NewCodeProcessingService() *CodeProcessingService {
	return &CodeProcessingService{}
}

func (s *CodeProcessingService) SubmitCode(ctx context.Context, req *pb.SubmitCodeReq) (*pb.SubmitCodeResp, error) {
	return &pb.SubmitCodeResp{}, nil
}
