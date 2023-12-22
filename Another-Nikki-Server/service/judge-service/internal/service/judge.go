package service

import (
	"Another-Nikki/service/judge-service/api"
	"context"
	"github.com/google/uuid"
)

func (s *server) Judge(ctx context.Context, req *api.JudgeReq) (resp *api.JudgeResp, err error) {
	resp = new(api.JudgeResp)
	ID := uuid.NewString()
	err = compile(ID, req.Code, "", req.Language)
	return
}

func (s *server) OnlineRun(ctx context.Context, req *api.OnlineRunReq) (resp *api.OnlineRunResp, err error) {
	resp = new(api.OnlineRunResp)
	ID := uuid.NewString()
	err = compile(ID, req.Code, req.Input, req.Language)
	return
}
