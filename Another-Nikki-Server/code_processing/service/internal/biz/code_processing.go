package biz

import (
	pb "Another-Nikki/code_processing/service/api"
	"golang.org/x/net/context"
)

type CodeDataRepo interface {
	CreateCode(ctx context.Context, req *pb.SubmitCodeReq) (err error)
}
