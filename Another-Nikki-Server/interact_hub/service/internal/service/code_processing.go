package service

import (
	"Another-Nikki/interact_hub/service/internal/biz"
	"Another-Nikki/interact_hub/service/internal/data"
	judge "Another-Nikki/judge/service/api"
	"Another-Nikki/pkg/log"
	"context"
	"fmt"
	"strconv"
	"time"

	pb "Another-Nikki/interact_hub/service/api"
)

type CodeProcessingService struct {
	pb.UnimplementedCodeProcessingServer

	dao         biz.CodeDataRepo
	judgeClient judge.JudgeClient
}

const maxCodeLength = 400 * 100

func NewCodeProcessingService(globalGrpc *data.GlobalGrpcClient, dao biz.CodeDataRepo) *CodeProcessingService {
	return &CodeProcessingService{
		dao:         dao,
		judgeClient: globalGrpc.JudgeClient,
	}
}

func (s *CodeProcessingService) judgeCode(ctx context.Context, req *pb.SubmitCodeReq, judgeId int64) {
	_, err := s.UpdateCodeCompileStatus(ctx, &pb.UpdateCodeCompileStatusReq{
		JudgeId:       judgeId,
		CompileStatus: "Compiling",
		CompileLog:    "",
	})
	if err != nil {
		log.Error(ctx, "modify compile status err: %v", err)
		return
	}

	var judgeResp *judge.JudgeResp
	for i := 1; i <= 3; i++ {
		judgeResp, err = s.judgeClient.Judge(ctx, &judge.JudgeReq{
			Code:        req.Code,
			Language:    judge.Language(judge.Language_value[req.Language]),
			ProblemName: strconv.Itoa(int(req.ProblemId)) + "." + req.ProblemName,
		})
		if err == nil {
			break
		}
	}
	if err != nil || judgeResp == nil {
		log.Error(ctx, "err = %v, judgeResp = %+v", err, judgeResp)
		return
	}
	if judgeResp.IsCompileError {
		_, _ = s.UpdateCodeCompileStatus(ctx, &pb.UpdateCodeCompileStatusReq{
			JudgeId:       judgeId,
			CompileStatus: judgeResp.CompileState,
			CompileLog:    judgeResp.CompileLog,
		})
		return
	}

	_, err = s.UpdateCodeJudgeStatus(ctx, &pb.UpdateCodeJudgeStatusReq{
		JudgeId:       judgeId,
		CompileStatus: judgeResp.CompileState,
		JudgeStatus:   judgeResp.JudgeResult,
		CpuTimeUsed:   judgeResp.CpuTimeUsed,
		MemoryUsed:    judgeResp.MemoryUsed,
	})
	if err != nil {
		log.Error(ctx, "modify judge resp err: %v", err)
		return
	}
}

func (s *CodeProcessingService) SubmitCode(ctx context.Context, req *pb.SubmitCodeReq) (resp *pb.SubmitCodeResp, err error) {
	resp = new(pb.SubmitCodeResp)
	if len(req.Code) > maxCodeLength {
		return nil, fmt.Errorf("代码过长, 请化简以后重新提交")
	}
	if len(req.Code) == 0 {
		return nil, fmt.Errorf("代码长度不可以为 0")
	}
	submitResp, err := s.dao.CreateCode(ctx, &biz.CreateCodeReq{
		UserId:      req.UserId,
		UserName:    req.UserName,
		ProblemId:   req.ProblemId,
		ProblemName: req.ProblemName,
		Language:    req.Language,
		Code:        req.Code,
	})
	resp.LastSubmitId = submitResp.LastInsertId
	s.judgeCode(ctx, req, submitResp.LastInsertId)
	return resp, err
}

func (s *CodeProcessingService) UpdateCodeCompileStatus(ctx context.Context, req *pb.UpdateCodeCompileStatusReq) (resp *pb.UpdateCodeCompileStatusResp, err error) {
	resp = new(pb.UpdateCodeCompileStatusResp)
	err = s.dao.UpdateCodeCompileStatus(ctx, &biz.UpdateCodeCompileStatusReq{
		CodeId:        req.JudgeId,
		CompileStatus: req.CompileStatus,
		CompileLog:    req.CompileLog,
	})
	return
}

func (s *CodeProcessingService) UpdateCodeJudgeStatus(ctx context.Context, req *pb.UpdateCodeJudgeStatusReq) (resp *pb.UpdateCodeJudgeStatusResp, err error) {
	resp = new(pb.UpdateCodeJudgeStatusResp)
	err = s.dao.UpdateCodeJudgeStatus(ctx, &biz.UpdateCodeJudgeStatusReq{
		CodeId:        req.JudgeId,
		CompileStatus: req.CompileStatus,
		JudgeStatus:   req.JudgeStatus,
		CpuTimeUsed:   req.CpuTimeUsed,
		MemoryUsed:    req.MemoryUsed,
	})
	return
}

func (s *CodeProcessingService) GetCommitByJudgeId(ctx context.Context, req *pb.GetCommitByJudgeIdReq) (resp *pb.GetCommitByJudgeIdResp, err error) {
	resp = new(pb.GetCommitByJudgeIdResp)
	res, err := s.dao.GetCommitByJudgeId(ctx, &biz.GetCommitByJudgeIdReq{JudgeId: req.JudgeId})
	if err != nil {
		return
	}
	resp.CreatedTime = res.CreatedTime.Format(time.DateTime)
	resp.Language = res.Language
	resp.ProblemId = res.ProblemId
	resp.JudgeStatus = res.JudgeStatus
	resp.Code = "```" + resp.Language + "\n" + res.Code + "\n```"
	resp.MemoryUsed = res.MemoryUsed
	resp.CpuTimeUsed = res.CpuTimeUsed
	resp.CompileStatus = res.CompileStatus
	resp.ProblemName = res.ProblemName
	resp.Username = res.UserName
	resp.UserId = res.UserId
	resp.CompileLog = res.CompileLog.String
	return
}
