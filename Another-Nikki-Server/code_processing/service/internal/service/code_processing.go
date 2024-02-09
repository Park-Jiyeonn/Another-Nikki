package service

import (
	pb "Another-Nikki/code_processing/service/api"
	"Another-Nikki/code_processing/service/internal/biz"
	"Another-Nikki/code_processing/service/internal/data"
	judge "Another-Nikki/judge/service/api"
	"context"
)

type CodeProcessingService struct {
	pb.UnimplementedCodeProcessingServer

	judgeClient judge.JudgeClient
	dao         *biz.CodeDataUseCase
}

func NewCodeProcessingService(globalGrpc *data.GlobalGrpcClient, dao *biz.CodeDataUseCase) *CodeProcessingService {
	return &CodeProcessingService{
		judgeClient: globalGrpc.JudgeClient,
		dao:         dao,
	}
}

func (s *CodeProcessingService) SubmitCode(ctx context.Context, req *pb.SubmitCodeReq) (resp *pb.SubmitCodeResp, err error) {
	resp = new(pb.SubmitCodeResp)
	err = s.dao.CreateCode(ctx, req)
	return resp, err
}
