package service

import (
	"context"

	pb "Another-Nikki/problem/service/api"
)

type ProblemService struct {
	pb.UnimplementedProblemServer
}

func NewProblemService() *ProblemService {
	return &ProblemService{}
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
