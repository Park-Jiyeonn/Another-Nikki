package service

import (
	"Another-Nikki/judge/service/api"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

type JudgeService struct {
	api.UnimplementedJudgeServer

	log *log.Helper
}

func NewJudgeService(logger log.Logger) *JudgeService {
	return &JudgeService{
		log: log.NewHelper(logger),
	}
}

func (s *JudgeService) Judge(ctx context.Context, req *api.JudgeReq) (resp *api.JudgeResp, err error) {
	//log.Info(ctx, "666666")
	ID := uuid.NewString()
	if err = compile(ID, req.Code, "", req.Language); err != nil {
		//s.log.WithContext(ctx).Error(err)
		return
	}
	if err = judge(ID, req.ProblemName); err != nil {
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
