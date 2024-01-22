package service

import (
	"Another-Nikki/service/judge/api"
	"Another-Nikki/util/log"
	"context"
	"github.com/google/uuid"
)

func (s *server) Judge(ctx context.Context, req *api.JudgeReq) (resp *api.JudgeResp, err error) {
	log.Info(ctx, "666666")
	ID := uuid.NewString()
	if err = compile(ID, req.Code, "", req.Language); err != nil {
		log.Error(ctx, "120210")
		return
	}
	if err = judge(ID, req.ProblemName); err != nil {
		return
	}
	resp, err = readJudgeRet(ID)
	if err != nil {
		log.Error(ctx, "")
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
