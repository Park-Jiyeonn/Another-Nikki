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
	"time"
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

func (s *UserService) GetUserCommitRecordByPage(ctx context.Context, req *pb.GetUserCommitRecordReq) (resp *pb.GetUserCommitRecordResp, err error) {
	resp = new(pb.GetUserCommitRecordResp)
	userId, _ := jwt.GetUserFromCtx(ctx)
	commits, err := s.dao.GetUserCommitRecord(ctx, &biz.GetUserCommitRecordReq{
		UserId:   userId,
		PageSize: req.PageSize,
		PageNum:  req.PageNum,
	})
	if err != nil {
		log.Error(ctx, "get commit record err: %v", err)
		return nil, err
	}
	for _, val := range commits {
		resp.Commits = append(resp.Commits, &pb.Commits{
			JudgeId:       val.JudgeId,
			ProblemName:   val.ProblemName,
			CompileStatus: val.CompileStatus,
			JudgeStatus:   val.JudgeStatus,
			CpuTimeUsed:   val.CpuTimeUsed,
			MemoryUsed:    val.MemoryUsed,
			Language:      val.Language,
			CreatedTime:   val.CreatedTime.Format(time.DateTime),
		})
	}
	return
}

func (s *UserService) GetUserSumCommit(ctx context.Context, _ *pb.GetUserSumCommitReq) (resp *pb.GetUserSumCommitResp, err error) {
	resp = new(pb.GetUserSumCommitResp)
	userId, _ := jwt.GetUserFromCtx(ctx)
	res, err := s.dao.GetUserSumCommit(ctx, &biz.GetUserSumCommitReq{userId})
	if err != nil {
		return
	}
	resp.Sum = res.Sum
	return
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (resp *pb.UpdateUserResp, err error) {
	resp = new(pb.UpdateUserResp)
	userId, _ := jwt.GetUserFromCtx(ctx)
	_, err = s.dao.UpdateUser(ctx, &biz.UpdateUserReq{
		UserId:   userId,
		Username: req.Username,
		Avatar:   req.Avatar,
	})
	return
}
