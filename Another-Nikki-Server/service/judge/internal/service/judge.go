package service

import (
	"Another-Nikki/service/judge/api"
	"context"
	"github.com/google/uuid"
)

func (s *server) Judge(ctx context.Context, req *api.JudgeReq) (resp *api.JudgeResp, err error) {
	ID := uuid.NewString()
	if err = compile(ID, req.Code, "", req.Language); err != nil {
		return
	}
	if err = judge(ID, req.ProblemName); err != nil {
		return
	}
	resp, err = readJudgeRet(ID)
	if err != nil {
		return
	}
	return
}

func (s *server) OnlineRun(ctx context.Context, req *api.OnlineRunReq) (resp *api.OnlineRunResp, err error) {
	ID := uuid.NewString()
	if err = compile(ID, req.Code, req.Input, req.Language); err != nil {
		return
	}
	if err = run(ID); err != nil {
		return
	}
	resp, err = readRunRet(ID)
	if err != nil {
		return
	}
	return
}
