package biz

import "golang.org/x/net/context"

type UserRepo interface {
	Register(ctx context.Context, req *RegisterReq) error
	GetUserByUserName(ctx context.Context, req *GetUserByUserNameReq) (*GetUserByUserNameResp, error)
	GetUserById(ctx context.Context, req *GetUserByIdReq) (*GetUserByIdResp, error)
}

type RegisterReq struct {
	Username string
	Password string
}

type GetUserByUserNameReq struct {
	Username string
}

type GetUserByUserNameResp struct {
	Username string
	Avatar   string
}

type GetUserByIdReq struct {
	UserId string
}

type GetUserByIdResp struct {
	Username string
	Avatar   string
}
