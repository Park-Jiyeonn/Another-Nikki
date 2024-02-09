package biz

import (
	pb "Another-Nikki/code_processing/service/api"
	"golang.org/x/net/context"
)

type CodeDataRepo interface {
	CreateCode(ctx context.Context, req *pb.SubmitCodeReq) (err error)
}

type CodeDataUseCase struct {
	repo CodeDataRepo
}

func NewCodeDataRepo(repo CodeDataRepo) *CodeDataUseCase {
	return &CodeDataUseCase{
		repo: repo,
	}
}

func (uc *CodeDataUseCase) CreateCode(ctx context.Context, req *pb.SubmitCodeReq) (err error) {
	err = uc.repo.CreateCode(ctx, req)
	return
}
