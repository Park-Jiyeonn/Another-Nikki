package service

import (
	pb "Another-Nikki/interact_hub/service/api"
	"Another-Nikki/interact_hub/service/internal/biz"
	"Another-Nikki/interact_hub/service/internal/conf"
	"Another-Nikki/pkg/jwt"
	"Another-Nikki/pkg/log"
	"context"
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strconv"
	"time"
)

type UserService struct {
	pb.UnimplementedUserServer

	dao    biz.UserRepo
	rx     *rand.Rand
	avatar []string
}

func NewUserService(dao biz.UserRepo, c *conf.Avatars) *UserService {
	return &UserService{
		dao:    dao,
		rx:     rand.New(rand.NewSource(time.Now().UnixNano())),
		avatar: c.Avatars,
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
	resp.Description = user.Description
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
		Username:    req.Username,
		Password:    string(newPassword),
		Avatar:      req.Avatar,
		Description: req.Description,
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
	resp.Description = user.Description
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
	resp.Description = user.Description
	return
}

func (s *UserService) GetUserCommitRecordByPage(ctx context.Context, req *pb.GetUserCommitRecordReq) (resp *pb.GetUserCommitRecordResp, err error) {
	resp = new(pb.GetUserCommitRecordResp)
	commits, err := s.dao.GetUserCommitRecord(ctx, &biz.GetUserCommitRecordReq{
		UserId:   req.UserId,
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
			ProblemId:     val.ProblemId,
		})
	}
	return
}

func (s *UserService) GetUserSumCommit(ctx context.Context, req *pb.GetUserSumCommitReq) (resp *pb.GetUserSumCommitResp, err error) {
	resp = new(pb.GetUserSumCommitResp)
	res, err := s.dao.GetUserSumCommit(ctx, &biz.GetUserSumCommitReq{UserId: req.UserId})
	if err != nil {
		return
	}
	resp.Sum = res.Sum
	return
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (resp *pb.UpdateUserResp, err error) {
	resp = new(pb.UpdateUserResp)
	userId, _ := jwt.GetUserFromCtx(ctx)
	user, err := s.dao.GetUserById(ctx, &biz.GetUserByIdReq{UserId: strconv.Itoa(int(userId))})
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("用户不存在")
	}
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword))
	if err != nil {
		return nil, fmt.Errorf("原密码错误，请重试")
	}
	newPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("修改失败，请换一个密码重试")
	}
	_, err = s.dao.UpdateUser(ctx, &biz.UpdateUserReq{
		UserId:      userId,
		Username:    req.Username,
		Avatar:      req.Avatar,
		Description: req.Description,
		Password:    string(newPassword),
	})
	return
}

func (s *UserService) CreateTouristAccount(ctx context.Context, _ *pb.CreateTouristAccountReq) (resp *pb.CreateTouristAccountResp, err error) {
	resp = new(pb.CreateTouristAccountResp)
	username := "user_"
	avatar := s.avatar[s.rx.Intn(len(s.avatar))]
	for i := 0; i < 8; i++ {
		username += string(rune(s.rx.Intn(10) + '0'))
	}
	res, err := s.Register(ctx, &pb.RegisterReq{
		Username:        username,
		Password:        "1234",
		ConfirmPassword: "1234",
		Avatar:          avatar,
		Description:     "lazy me, no description",
	})
	if err != nil {
		log.Error(ctx, "CreateTouristAccount err: %v", err)
		return
	}
	resp.Token = res.Token
	resp.Username = username
	resp.UserId = res.UserId
	resp.Avatar = avatar
	resp.Description = "lazy me, no description"
	return
}

func (s *UserService) GetUserWrongProblem(ctx context.Context, req *pb.GetUserWrongProblemReq) (resp *pb.GetUserWrongProblemResp, err error) {
	resp = new(pb.GetUserWrongProblemResp)
	userWrongAnswerProblems, err := s.dao.GetUserWrongProblem(ctx, &biz.GetUserWrongProblemReq{UserId: req.UserId})
	if err != nil {
		return
	}
	for _, val := range userWrongAnswerProblems {
		resp.UserWrongProblem = append(resp.UserWrongProblem, &pb.UserWrongProblem{
			ProblemId:   val.ProblemId,
			ProblemName: val.ProblemName,
		})
	}
	return
}
