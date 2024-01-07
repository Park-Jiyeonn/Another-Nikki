package service

import (
	"Another-Nikki/service/code-processing/api"
	"context"
)

func (s *server) SubmitCode(ctx context.Context, req *api.SubmitCodeReq) (resp *api.SubmitCodeResp, err error) {
	resp = new(api.SubmitCodeResp)
	return
}
