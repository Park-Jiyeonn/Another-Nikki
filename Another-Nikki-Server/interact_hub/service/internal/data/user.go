package data

import (
	"Another-Nikki/interact_hub/service/internal/biz"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
)

type userServiceImpl struct {
	db *sqlx.DB
}

// CREATE TABLE users (
//    user_id SERIAL PRIMARY KEY,
//    username VARCHAR(255) UNIQUE NOT NULL,
//    password VARCHAR(255) NOT NULL,
//    avatar VARCHAR(255),
//    created_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//    updated_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
// );
//
//-- Create an index on the 'username' column
// CREATE INDEX idx_username ON users (username);

func NewUserImpl(data *Data) biz.UserRepo {
	return &userServiceImpl{
		data.GlobalDB,
	}
}

func (s *userServiceImpl) Register(ctx context.Context, req *biz.RegisterReq) (resp *biz.RegisterResp, err error) {
	ret, err := s.db.ExecContext(ctx, "INSERT INTO users (username, password) VALUES (?, ?)", req.Username, req.Password)
	if err != nil {
		return
	}
	resp = new(biz.RegisterResp)
	resp.UserId, err = ret.LastInsertId()
	return
}

func (s *userServiceImpl) GetUserByUserName(ctx context.Context, req *biz.GetUserByUserNameReq) (*biz.GetUserByUserNameResp, error) {
	var username, avatar, password string
	err := s.db.QueryRowContext(ctx, "SELECT username, avatar, password FROM users WHERE username = ?", req.Username).Scan(&username, &avatar, &password)
	if err != nil {
		return nil, err
	}
	return &biz.GetUserByUserNameResp{Username: username, Avatar: avatar}, nil
}

func (s *userServiceImpl) GetUserById(ctx context.Context, req *biz.GetUserByIdReq) (*biz.GetUserByIdResp, error) {
	var username, avatar string
	err := s.db.QueryRowContext(ctx, "SELECT username, avatar FROM users WHERE user_id = ?", req.UserId).Scan(&username, &avatar)
	if err != nil {
		return nil, err
	}
	return &biz.GetUserByIdResp{Username: username, Avatar: avatar}, nil
}
