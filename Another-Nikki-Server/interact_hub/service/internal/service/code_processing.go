package service

import (
	"Another-Nikki/interact_hub/service/internal/biz"
	"Another-Nikki/interact_hub/service/internal/data"
	"context"
	"fmt"

	pb "Another-Nikki/interact_hub/service/api"
)

type CodeProcessingService struct {
	pb.UnimplementedCodeProcessingServer

	dao biz.CodeDataRepo
}

const maxCodeLength = 400 * 100

func NewCodeProcessingService(globalGrpc *data.GlobalGrpcClient, dao biz.CodeDataRepo) *CodeProcessingService {
	return &CodeProcessingService{
		dao: dao,
	}
}

func (s *CodeProcessingService) SubmitCode(ctx context.Context, req *pb.SubmitCodeReq) (resp *pb.SubmitCodeResp, err error) {
	resp = new(pb.SubmitCodeResp)
	if len(req.Code) > maxCodeLength {
		return nil, fmt.Errorf("代码过长, 请化简以后重新提交")
	}
	err = s.dao.CreateCode(ctx, &biz.CreateCodeReq{
		UserId:      req.UserId,
		UserName:    req.UserName,
		ProblemId:   req.ProblemId,
		ProblemName: req.ProblemName,
		Language:    req.Language,
		Code:        req.Code,
	})
	return resp, err
}

func (s *CodeProcessingService) UpdateCodeCompileStatus(ctx context.Context, req *pb.UpdateCodeCompileStatusReq) (resp *pb.UpdateCodeCompileStatusResp, err error) {
	resp = new(pb.UpdateCodeCompileStatusResp)
	err = s.dao.UpdateCodeCompileStatus(ctx, &biz.UpdateCodeCompileStatusReq{
		CodeId:        req.CodeId,
		CompileStatus: req.CompileStatus,
		CompileLog:    req.CompileLog,
	})
	return
}

func (s *CodeProcessingService) UpdateCodeJudgeStatus(ctx context.Context, req *pb.UpdateCodeJudgeStatusReq) (resp *pb.UpdateCodeJudgeStatusResp, err error) {
	resp = new(pb.UpdateCodeJudgeStatusResp)
	err = s.dao.UpdateCodeJudgeStatus(ctx, &biz.UpdateCodeJudgeStatusReq{
		CodeId:        req.CodeId,
		CompileStatus: req.CompileStatus,
		JudgeStatus:   req.JudgeStatus,
		CpuTimeUsed:   req.CpuTimeUsed,
		MemoryUsed:    req.MemoryUsed,
	})
	return
}
