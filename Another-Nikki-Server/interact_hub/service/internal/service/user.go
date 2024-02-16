package service

import (
	pb "Another-Nikki/interact_hub/service/api"
	"Another-Nikki/interact_hub/service/internal/biz"
	"Another-Nikki/pkg/jwt"
	"Another-Nikki/pkg/log"
	"context"
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
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

func (s *UserService) Login(ctx context.Context, req *pb.LoginReq) (resp *pb.LoginResp, err error) {
	resp = new(pb.LoginResp)
	if req.Username == "" || req.Password == "" {
		return nil, fmt.Errorf("用户名或密码为空")
	}
	if len(req.Username) > 20 || len(req.Password) > 20 {
		return nil, fmt.Errorf("用户名或密码太长，不可以超过 20")
	}
	user, err := s.dao.GetUserByUserName(ctx, &biz.GetUserByUserNameReq{Username: req.Username})
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("用户不存在")
	}
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, fmt.Errorf("登录失败，请重试")
	}
	resp.Token, err = jwt.GenToken(user.UserId, req.Username)
	if err != nil {
		log.Error(ctx, "gen jwt token err: %v", err)
		return
	}
	resp.Username = req.Username
	resp.UserId = user.UserId
	resp.Avatar = user.Avatar
	return
}
func (s *UserService) Register(ctx context.Context, req *pb.RegisterReq) (resp *pb.RegisterResp, err error) {
	if req.Username == "" || req.Password == "" {
		return nil, fmt.Errorf("用户名或密码为空")
	}
	if len(req.Username) > 20 || len(req.Password) > 20 || len(req.ConfirmPassword) > 20 {
		return nil, fmt.Errorf("用户名或密码太长，不可以超过 20")
	}
	if req.Password != req.ConfirmPassword {
		return nil, fmt.Errorf("密码不一致捏")
	}
	_, err = s.dao.GetUserByUserName(ctx, &biz.GetUserByUserNameReq{Username: req.Username})
	if err == nil {
		return nil, fmt.Errorf("用户名已经存在")
	}

	newPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("注册失败，请换一个密码重试")
	}
	user, err := s.dao.Register(ctx, &biz.RegisterReq{
		Username: req.Username,
		Password: string(newPassword),
	})
	if err != nil {
		return nil, fmt.Errorf("注册失败，请重试")
	}

	resp = new(pb.RegisterResp)
	resp.Token, err = jwt.GenToken(user.UserId, req.Username)
	if err != nil {
		log.Error(ctx, "gen jwt token err: %v", err)
		return
	}
	resp.Username = req.Username
	resp.UserId = user.UserId
	return
}
func (s *UserService) GetUserByUserName(ctx context.Context, req *pb.GetUserByUserNameReq) (resp *pb.GetUserByUserNameResp, err error) {
	resp = new(pb.GetUserByUserNameResp)
	user, err := s.dao.GetUserByUserName(ctx, &biz.GetUserByUserNameReq{Username: req.Username})
	if err != nil {
		return nil, err
	}
	resp.Username = user.Username
	resp.Avatar = user.Avatar
	return
}
func (s *UserService) GetUserById(ctx context.Context, req *pb.GetUserByIdReq) (resp *pb.GetUserByIdResp, err error) {
	resp = new(pb.GetUserByIdResp)
	user, err := s.dao.GetUserById(ctx, &biz.GetUserByIdReq{UserId: req.UserId})
	if err != nil {
		return nil, err
	}
	resp.Username = user.Username
	resp.Avatar = user.Avatar
	return
}
