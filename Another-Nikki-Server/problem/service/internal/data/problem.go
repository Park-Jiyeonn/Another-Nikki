package data

import (
	"Another-Nikki/problem/service/internal/biz"
	"github.com/jmoiron/sqlx"
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

// CREATE TABLE judge (
//    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
//    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
//    user_id BIGINT NOT NULL DEFAULT 0,
//    user_name VARCHAR(255) NOT NULL DEFAULT '',
//    problem_id BIGINT NOT NULL DEFAULT 0,
//    problem_name VARCHAR(255) NOT NULL DEFAULT '',
//    language VARCHAR(255) NOT NULL DEFAULT '',
//    code TEXT NOT NULL,
//    compile_status VARCHAR(255) NOT NULL DEFAULT 'in queue',
//    compile_log VARCHAR(255) NOT NULL DEFAULT '',
//    judge_status VARCHAR(255) NOT NULL DEFAULT '-',
//    cpu_time_used VARCHAR(255) NOT NULL DEFAULT '0',
//    memory_used VARCHAR(255) NOT NULL DEFAULT '0'
// );

func (s *problemRepo) CreateProblem(problemId int64, title, description, content string) (err error) {
	return
}
func (s *problemRepo) GetProblemById(problemId int64) (problemTitle, problemDescription, problemContent, createTime string, err error) {
	return
}
func (s *problemRepo) GetProblemByPage(pageNum, pageSize int64) (problems []*biz.Problems, err error) {
	return
}
