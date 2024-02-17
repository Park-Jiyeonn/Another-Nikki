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
	resp = new(api.JudgeResp)
	ID := uuid.NewString()
	if resp.CompileLog, resp.IsCompileError = compile(ctx, ID, req.Code, "", req.Language); resp.IsCompileError {
		resp.CompileState = "Compile Failed"
		return
	}
	resp.CompileState = "Compile Success"
	if err = judge(ctx, ID, req.ProblemName, req.Language); err != nil {
		return
	}
	if req.Language == api.Language_Python {
		ret, ok := checkPythonSyntaxError(ID)
		if ok {
			resp.CompileState = "Compile Failed"
			resp.CompileLog = ret
			resp.IsCompileError = true
			return
		}
	}
	resp.MemoryUsed, resp.CpuTimeUsed, resp.JudgeResult, err = readJudgeRet(ID)
	if err != nil {
		return
	}
	return
}

func (s *JudgeService) OnlineRun(ctx context.Context, req *api.OnlineRunReq) (resp *api.OnlineRunResp, err error) {
	ID := uuid.NewString()
	resp = new(api.OnlineRunResp)
	if resp.CompileLog, resp.IsCompileError = compile(ctx, ID, req.Code, req.Input, req.Language); resp.IsCompileError {
		resp.CompileState = "Compile Failed"
		return
	}
	resp.CompileState = "Compile Success"
	if err = run(ctx, ID, req.Language); err != nil {
		return
	}
	if req.Language == api.Language_Python {
		ret, ok := checkPythonSyntaxError(ID)
		if ok {
			resp.CompileState = "Compile Failed"
			resp.CompileLog = ret
			resp.IsCompileError = true
			return
		}
	}
	resp.Output, resp.MemoryUsed, resp.CpuTimeUsed, err = readRunRet(ID)
	if err != nil {
		return
	}
	return
}
