package biz

import (
	pb "Another-Nikki/code_processing/service/api"
	"golang.org/x/net/context"
)

type CodeDataRepo interface {
	CreateCode(ctx context.Context, req *pb.SubmitCodeReq) (err error)
	UpdateCodeCompileStatus(ctx context.Context, id int64, status, compile_log string) (err error)
	UpdateCodeJudgeStatus(ctx context.Context, id int64, status string) (err error)
}
