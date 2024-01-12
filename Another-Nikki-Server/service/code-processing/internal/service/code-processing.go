package service

import (
	"Another-Nikki/dal"
	"Another-Nikki/service/code-processing/api"
	"context"
)

var judge dal.Judge

func (s *server) SubmitCode(ctx context.Context, req *api.SubmitCodeReq) (resp *api.SubmitCodeResp, err error) {
	resp = new(api.SubmitCodeResp)
	err = judge.Create(ctx, req.UserID, req.UserName, req.ProblemID, req.ProblemName, req.Language, req.Code)
	return
}
