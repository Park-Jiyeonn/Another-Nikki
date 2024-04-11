package biz

import (
	"golang.org/x/net/context"
	"time"
)

type UserRepo interface {
	Register(ctx context.Context, req *RegisterReq) (*RegisterResp, error)
	GetUserByUserName(ctx context.Context, req *GetUserByUserNameReq) (*GetUserByUserNameResp, error)
	GetUserById(ctx context.Context, req *GetUserByIdReq) (*GetUserByIdResp, error)
	GetUserCommitRecord(ctx context.Context, req *GetUserCommitRecordReq) (resp []*GetUserCommitRecordResp, err error)
	GetUserSumCommit(ctx context.Context, req *GetUserSumCommitReq) (resp *GetUserSumCommitResp, err error)
	UpdateUser(ctx context.Context, req *UpdateUserReq) (resp *UpdateUserResp, err error)
}

type RegisterReq struct {
	Username    string
	Password    string
	Avatar      string
	Description string
}

type RegisterResp struct {
	UserId int64 `db:"user_id"`
}

type GetUserByUserNameReq struct {
	Username string
}

type GetUserByUserNameResp struct {
	Username    string
	Password    string
	Avatar      string
	UserId      int64
	Description string
}

type GetUserByIdReq struct {
	UserId string
}

type GetUserByIdResp struct {
	Username    string
	Password    string
	Avatar      string
	UserId      int64
	Description string
}

type GetUserCommitRecordReq struct {
	UserId   int64
	PageNum  int64
	PageSize int64
}

type GetUserCommitRecordResp struct {
	JudgeId       int64     `db:"judge_id"`
	ProblemId     int64     `db:"problem_id"`
	ProblemName   string    `db:"problem_name"`
	CompileStatus string    `db:"compile_status"`
	JudgeStatus   string    `db:"judge_status"`
	CpuTimeUsed   string    `db:"cpu_time_used"`
	MemoryUsed    string    `db:"memory_used"`
	Language      string    `db:"language"`
	CreatedTime   time.Time `db:"created_time"`
}

type GetUserSumCommitReq struct {
	UserId int64
}
type GetUserSumCommitResp struct {
	Sum int64
}

type UpdateUserReq struct {
	UserId      int64
	Username    string
	Avatar      string
	Description string
	Password    string
}
type UpdateUserResp struct {
}
