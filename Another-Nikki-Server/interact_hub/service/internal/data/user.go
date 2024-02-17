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
//	  description VARCHAR(255) DEFAULT '这个人很懒，没有个性签名',
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
	var username, avatar, password, description string
	var user_id int64
	err := s.db.QueryRowContext(ctx, "SELECT username, avatar, password, user_id, description FROM users WHERE username = ?", req.Username).Scan(&username, &avatar, &password, &user_id, &description)
	if err != nil {
		return nil, err
	}
	return &biz.GetUserByUserNameResp{Username: username,
		Avatar:      avatar,
		UserId:      user_id,
		Password:    password,
		Description: description,
	}, nil
}

func (s *userServiceImpl) GetUserById(ctx context.Context, req *biz.GetUserByIdReq) (*biz.GetUserByIdResp, error) {
	var username, avatar, description string
	err := s.db.QueryRowContext(ctx, "SELECT username, avatar, description FROM users WHERE user_id = ?", req.UserId).Scan(&username, &avatar, &description)
	if err != nil {
		return nil, err
	}
	return &biz.GetUserByIdResp{
		Username:    username,
		Avatar:      avatar,
		Description: description,
	}, nil
}

func (s *userServiceImpl) GetUserCommitRecord(ctx context.Context, req *biz.GetUserCommitRecordReq) (resp []*biz.GetUserCommitRecordResp, err error) {
	sqlStr := "SELECT judge_id, problem_id, problem_name, compile_status, judge_status, cpu_time_used, memory_used, language, created_time from judges where user_id = ? ORDER BY created_time DESC LIMIT ? OFFSET ?"
	rows, err := s.db.QueryxContext(ctx, sqlStr, req.UserId, req.PageSize, (req.PageNum-1)*req.PageSize)
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

func (s *userServiceImpl) GetUserSumCommit(ctx context.Context, req *biz.GetUserSumCommitReq) (resp *biz.GetUserSumCommitResp, err error) {
	resp = new(biz.GetUserSumCommitResp)
	err = s.db.QueryRowContext(ctx, "SELECT COUNT(judge_id) from judges where user_id = ?", req.UserId).Scan(&resp.Sum)
	return
}

func (s *userServiceImpl) UpdateUser(ctx context.Context, req *biz.UpdateUserReq) (resp *biz.UpdateUserResp, err error) {
	resp = new(biz.UpdateUserResp)
	sqlStr := "UPDATE users set username = ?, avatar = ?, description = ? where user_id = ?"
	_, err = s.db.ExecContext(ctx, sqlStr, req.Username, req.Avatar, req.Description, req.UserId)
	return
}
