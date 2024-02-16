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
//    avatar VARCHAR(255) DEFAULT '',
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
	var user_id int64
	err := s.db.QueryRowContext(ctx, "SELECT username, avatar, password, user_id FROM users WHERE username = ?", req.Username).Scan(&username, &avatar, &password, &user_id)
	if err != nil {
		return nil, err
	}
	return &biz.GetUserByUserNameResp{Username: username,
		Avatar:   avatar,
		UserId:   user_id,
		Password: password,
	}, nil
}

func (s *userServiceImpl) GetUserById(ctx context.Context, req *biz.GetUserByIdReq) (*biz.GetUserByIdResp, error) {
	var username, avatar string
	err := s.db.QueryRowContext(ctx, "SELECT username, avatar FROM users WHERE user_id = ?", req.UserId).Scan(&username, &avatar)
	if err != nil {
		return nil, err
	}
	return &biz.GetUserByIdResp{
		Username: username,
		Avatar:   avatar,
	}, nil
}

func (s *userServiceImpl) GetUserCommitRecord(ctx context.Context, req *biz.GetUserCommitRecordReq) (resp []*biz.GetUserCommitRecordResp, err error) {
	sqlStr := "SELECT judge_id, problem_name, compile_status, judge_status, cpu_time_used, memory_used, language, created_time from judges where user_id = ?"
	rows, err := s.db.QueryxContext(ctx, sqlStr, req.UserId)
	if err != nil {
		return nil, err
	}
	resp = make([]*biz.GetUserCommitRecordResp, 0)
	for rows.Next() {
		var res biz.GetUserCommitRecordResp
		err = rows.StructScan(&res)
		resp = append(resp, &res)
		if err != nil {
			return
		}
	}
	return
}
