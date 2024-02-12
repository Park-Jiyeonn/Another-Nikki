package service

import (
	"Another-Nikki/interact_hub/service/internal/biz"
	"context"

	pb "Another-Nikki/interact_hub/service/api"
)

type ProblemService struct {
	pb.UnimplementedProblemServer
	dao biz.ProblemRepo
}

func NewProblemService(dao biz.ProblemRepo) *ProblemService {
	return &ProblemService{
		dao: dao,
	}
}

func (s *ProblemService) PostProblem(ctx context.Context, req *pb.PostProblemReq) (*pb.PostProblemResp, error) {
	return &pb.PostProblemResp{}, nil
}
func (s *ProblemService) GetProblemById(ctx context.Context, req *pb.GetProblemByIdReq) (*pb.GetProblemByIdResp, error) {
	return &pb.GetProblemByIdResp{}, nil
}
func (s *ProblemService) GetProblemByPage(ctx context.Context, req *pb.GetProblemByPageReq) (*pb.GetProblemByPageResp, error) {
	return &pb.GetProblemByPageResp{}, nil
}
