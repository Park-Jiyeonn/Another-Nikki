package data

import (
	"Another-Nikki/code_processing/service/api"
	"Another-Nikki/code_processing/service/internal/biz"
	"Another-Nikki/pkg/log"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
	"time"
)

type CodeDataMysql struct {
	ID            int64 `db:"id"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	UserID        int64  `db:"user_id"`
	UserName      string `db:"user_name"`
	ProblemID     int64  `db:"problem_id"`
	ProblemName   string `db:"problem_name"`
	Language      string `db:"language"`
	Code          string `db:"code"`
	CompileStatus string `db:"compile_status"`
	JudgeStatus   string `db:"judge_status"`
}

type codeProcessingRepo struct {
	db *sqlx.DB
}

func (c *codeProcessingRepo) CreateCode(ctx context.Context, req *api.SubmitCodeReq) (err error) {
	_, err = c.db.ExecContext(ctx, "INSERT INTO judge (user_id, user_name, problem_id, problem_name, language, code) VALUES (?, ?, ?, ?, ?, ?)",
		req.UserID,
		req.UserName,
		req.ProblemID,
		req.ProblemName,
		req.Language,
		req.Code,
	)
	if err != nil {
		log.Error(ctx, "create judge code err: %v", err)
	}
	return
}

func NewCodeProcessingRepo(data *Data) biz.CodeDataRepo {
	return &codeProcessingRepo{
		db: data.GlobalDB,
	}
}

// CREATE TABLE judge (
//    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
//    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
//    user_id BIGINT NOT NULL DEFAULT '0',
//    user_name VARCHAR(255) NOT NULL DEFAULT '',
//    problem_id BIGINT NOT NULL DEFAULT '0',
//    problem_name VARCHAR(255) NOT NULL DEFAULT '',
//    language VARCHAR(255) NOT NULL DEFAULT '',
//    code TEXT NOT NULL,
//    compile_status VARCHAR(255) NOT NULL DEFAULT 'Compiling...',
//    judge_status VARCHAR(255) NOT NULL DEFAULT '-'
//);
