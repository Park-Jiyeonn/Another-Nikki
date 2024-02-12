package service

import (
	"Another-Nikki/interact_hub/service/internal/biz"
	"context"

	pb "Another-Nikki/interact_hub/service/api"
)

type UserService struct {
	pb.UnimplementedUserServer

	dao biz.UserRepo
}

func NewUserService(dao biz.UserRepo) *UserService {
	return &UserService{
		dao: dao,
	}
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	return &pb.LoginResp{}, nil
}
func (s *UserService) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	return &pb.RegisterResp{}, nil
}
func (s *UserService) GetUserByUserName(ctx context.Context, req *pb.GetUserByUserNameReq) (*pb.GetUserByUserNameResp, error) {
	return &pb.GetUserByUserNameResp{}, nil
}
func (s *UserService) GetUserById(ctx context.Context, req *pb.GetUserByIdReq) (*pb.GetUserByIdResp, error) {
	return &pb.GetUserByIdResp{}, nil
}
