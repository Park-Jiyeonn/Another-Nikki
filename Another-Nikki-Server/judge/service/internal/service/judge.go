package service

import (
	"Another-Nikki/judge/service/api"
	"context"
	"github.com/google/uuid"
)

type JudgeService struct {
	api.UnimplementedJudgeServer
}

func NewJudgeService() *JudgeService {
	return &JudgeService{}
}

func (s *JudgeService) Judge(ctx context.Context, req *api.JudgeReq) (resp *api.JudgeResp, err error) {
	ID := uuid.NewString()
	if err = compile(ctx, ID, req.Code, "", req.Language); err != nil {
		return
	}
	if err = judge(ctx, ID, req.ProblemName, req.Language); err != nil {
		return
	}
	resp, err = readJudgeRet(ID)
	if err != nil {
		//log.Error(ctx, "")
		return
	}
	return
}

func (s *JudgeService) OnlineRun(ctx context.Context, req *api.OnlineRunReq) (resp *api.OnlineRunResp, err error) {
	ID := uuid.NewString()
	if err = compile(ctx, ID, req.Code, req.Input, req.Language); err != nil {
		return
	}
	if err = run(ctx, ID, req.Language); err != nil {
		return
	}
	resp, err = readRunRet(ID)
	if err != nil {
		return
	}
	return
}
