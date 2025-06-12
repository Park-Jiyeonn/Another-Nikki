package data_sqlite

import (
	"Another-Nikki/interact_hub/service/internal/biz"
	"Another-Nikki/pkg/log"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
)

type userServiceImpl struct {
	db *sqlx.DB
}

func NewUserImpl(data *Data) biz.UserRepo {
	if err := InitUsersTable(data.GlobalDB); err != nil {
		panic(err)
	}
	return &userServiceImpl{
		data.GlobalDB,
	}
}

func InitUsersTable(db *sqlx.DB) error {
	const schema = `
	CREATE TABLE IF NOT EXISTS users (
		user_id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL,
		avatar TEXT DEFAULT '',
		description TEXT DEFAULT '这个人很懒，没有个性签名',
		created_time DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_time DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	CREATE INDEX IF NOT EXISTS idx_username ON users (username);
	`
	_, err := db.Exec(schema)
	return err
}

// 注册用户
func (s *userServiceImpl) Register(ctx context.Context, req *biz.RegisterReq) (resp *biz.RegisterResp, err error) {
	ret, err := s.db.ExecContext(ctx,
		`INSERT INTO users (username, password, avatar, description, updated_time)
		 VALUES (?, ?, ?, ?, CURRENT_TIMESTAMP)`,
		req.Username, req.Password, req.Avatar, req.Description)
	if err != nil {
		return
	}
	resp = new(biz.RegisterResp)
	resp.UserId, err = ret.LastInsertId()
	return
}

// 通过用户名查询用户
func (s *userServiceImpl) GetUserByUserName(ctx context.Context, req *biz.GetUserByUserNameReq) (*biz.GetUserByUserNameResp, error) {
	var username, avatar, password, description string
	var user_id int64
	err := s.db.QueryRowContext(ctx,
		`SELECT username, avatar, password, user_id, description 
		 FROM users 
		 WHERE username = ?`, req.Username).Scan(&username, &avatar, &password, &user_id, &description)
	if err != nil {
		return nil, err
	}
	return &biz.GetUserByUserNameResp{
		Username:    username,
		Avatar:      avatar,
		UserId:      user_id,
		Password:    password,
		Description: description,
	}, nil
}

// 通过用户 ID 查询
func (s *userServiceImpl) GetUserById(ctx context.Context, req *biz.GetUserByIdReq) (*biz.GetUserByIdResp, error) {
	var username, avatar, description, password string
	err := s.db.QueryRowContext(ctx,
		`SELECT username, avatar, description, password 
		 FROM users 
		 WHERE user_id = ?`, req.UserId).Scan(&username, &avatar, &description, &password)
	if err != nil {
		return nil, err
	}
	return &biz.GetUserByIdResp{
		Username:    username,
		Avatar:      avatar,
		Description: description,
		Password:    password,
	}, nil
}

// 获取用户提交记录
func (s *userServiceImpl) GetUserCommitRecord(ctx context.Context, req *biz.GetUserCommitRecordReq) (resp []*biz.GetUserCommitRecordResp, err error) {
	sqlStr := `SELECT judge_id, problem_id, problem_name, compile_status, judge_status, cpu_time_used, memory_used, language, created_time 
			   FROM judges 
			   WHERE user_id = ? 
			   ORDER BY created_time DESC 
			   LIMIT ? OFFSET ?`
	rows, err := s.db.QueryxContext(ctx, sqlStr, req.UserId, req.PageSize, (req.PageNum-1)*req.PageSize)
	if err != nil {
		return nil, err
	}
	resp = make([]*biz.GetUserCommitRecordResp, 0)
	for rows.Next() {
		var res biz.GetUserCommitRecordResp
		err = rows.StructScan(&res)
		if err != nil {
			return
		}
		resp = append(resp, &res)
	}
	return
}

// 获取用户总提交数
func (s *userServiceImpl) GetUserSumCommit(ctx context.Context, req *biz.GetUserSumCommitReq) (resp *biz.GetUserSumCommitResp, err error) {
	resp = new(biz.GetUserSumCommitResp)
	err = s.db.QueryRowContext(ctx,
		`SELECT COUNT(judge_id) 
		 FROM judges 
		 WHERE user_id = ?`, req.UserId).Scan(&resp.Sum)
	return
}

// 更新用户信息
func (s *userServiceImpl) UpdateUser(ctx context.Context, req *biz.UpdateUserReq) (resp *biz.UpdateUserResp, err error) {
	resp = new(biz.UpdateUserResp)
	sqlStr := `UPDATE users 
			   SET username = ?, avatar = ?, description = ?, password = ?, updated_time = CURRENT_TIMESTAMP 
			   WHERE user_id = ?`
	_, err = s.db.ExecContext(ctx, sqlStr,
		req.Username, req.Avatar, req.Description, req.Password, req.UserId)
	return
}

// 获取用户错误题目
func (s *userServiceImpl) GetUserWrongProblem(ctx context.Context, req *biz.GetUserWrongProblemReq) (resp []*biz.GetUserWrongProblemResp, err error) {
	sqlStr := `SELECT problem_id, problem_name 
			   FROM judges 
			   WHERE user_id = ?`
	rows, err := s.db.QueryxContext(ctx, sqlStr, req.UserId)
	if err != nil {
		return nil, err
	}

	resp = make([]*biz.GetUserWrongProblemResp, 0)
	f := make(map[int64]struct{})

	for rows.Next() {
		var res biz.GetUserWrongProblemResp
		err = rows.StructScan(&res)
		if err != nil {
			log.Error(ctx, "StructScan from GetUserWrongProblemResp err: %v", err)
			continue
		}

		// 查询是否有 AC 的记录
		queAcStr := `SELECT compile_status 
				     FROM judges 
				     WHERE user_id = ? AND problem_id = ? AND judge_status = ?`
		acRows, _ := s.db.QueryxContext(ctx, queAcStr, req.UserId, res.ProblemId, "Accept")
		if err != nil || acRows.Next() {
			continue
		}

		if _, ok := f[res.ProblemId]; ok {
			continue
		}
		f[res.ProblemId] = struct{}{}
		resp = append(resp, &res)
	}
	return
}
