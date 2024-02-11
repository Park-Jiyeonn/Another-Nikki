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
	dao         biz.CodeDataRepo
}

func NewCodeProcessingService(globalGrpc *data.GlobalGrpcClient, dao biz.CodeDataRepo) *CodeProcessingService {
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

func (s *CodeProcessingService) UpdateCodeCompileStatus(ctx context.Context, req *pb.UpdateCodeCompileStatusReq) (resp *pb.UpdateCodeCompileStatusResp, err error) {
	resp = new(pb.UpdateCodeCompileStatusResp)
	err = s.dao.UpdateCodeCompileStatus(ctx, req.CodeId, req.Status, req.CompileLog)
	return
}

func (s *CodeProcessingService) UpdateCodeJudgeStatus(ctx context.Context, req *pb.UpdateCodeJudgeStatusReq) (resp *pb.UpdateCodeJudgeStatusResp, err error) {
	resp = new(pb.UpdateCodeJudgeStatusResp)
	err = s.dao.UpdateCodeJudgeStatus(ctx, req.CodeId, req.Status)
	return
}
