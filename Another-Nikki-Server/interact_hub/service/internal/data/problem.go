package data

import (
	"Another-Nikki/interact_hub/service/internal/biz"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
	"time"
)

type ProblemDataMysql struct {
	ProblemId int64 `db:"problem_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type problemRepo struct {
	db *sqlx.DB
}

func NewProblemRepo(data *Data) biz.ProblemRepo {
	return &problemRepo{
		db: data.GlobalDB,
	}
}

// CREATE TABLE problems (
//    problem_id SERIAL PRIMARY KEY,
//    problem_title VARCHAR(255) NOT NULL,
//    problem_description VARCHAR(255),
//    problem_content TEXT,
//    created_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//    updated_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
// );

func (s *problemRepo) PostProblem(ctx context.Context, req *biz.PostProblemReq) (err error) {
	_, err = s.db.ExecContext(ctx, "INSERT INTO problems (problem_title, problem_description, problem_content) VALUES (?, ?, ?)",
		req.ProblemTitle, req.ProblemDescription, req.ProblemContent)
	return
}
func (s *problemRepo) GetProblemById(ctx context.Context, req *biz.GetProblemByIdReq) (resp *biz.GetProblemByIdResp, err error) {
	resp = new(biz.GetProblemByIdResp)
	err = s.db.GetContext(ctx, resp, "SELECT problem_title, problem_description, problem_content, created_time FROM problems WHERE problem_id = ?", req.ProblemId)
	return
}
func (s *problemRepo) GetProblemByPage(ctx context.Context, req *biz.GetProblemByPageReq) (resp *biz.GetProblemByPageResp, err error) {
	rows, err := s.db.QueryxContext(ctx, "SELECT problem_id, problem_title, created_time FROM problems ORDER BY created_time DESC LIMIT ? OFFSET ?",
		req.PageSize, (req.PageNum-1)*req.PageSize)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var problem biz.ProblemPageDetail
		err = rows.StructScan(&problem)
		if err != nil {
			return nil, err
		}
		resp.Problems = append(resp.Problems, &problem)
	}
	return
}
