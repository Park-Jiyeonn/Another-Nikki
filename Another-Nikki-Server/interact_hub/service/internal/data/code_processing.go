package data

import (
	"Another-Nikki/interact_hub/service/internal/biz"
	"Another-Nikki/pkg/log"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
)

type codeProcessingImpl struct {
	db *sqlx.DB
}

func NewCodeProcessingImpl(data *Data) biz.CodeDataRepo {
	return &codeProcessingImpl{
		db: data.GlobalDB,
	}
}

// CREATE TABLE judges (
//    judge_id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
//    created_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//    updated_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
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

func (c *codeProcessingImpl) CreateCode(ctx context.Context, req *biz.CreateCodeReq) (err error) {
	const sqlStr = "INSERT INTO judges (user_id, user_name, problem_id, problem_name, language, code) VALUES (?, ?, ?, ?, ?, ?)"
	_, err = c.db.ExecContext(ctx, sqlStr,
		req.UserId,
		req.UserName,
		req.ProblemId,
		req.ProblemName,
		req.Language,
		req.Code,
	)
	if err != nil {
		log.Error(ctx, "create judge code err: %v", err)
	}
	return
}

func (c *codeProcessingImpl) UpdateCodeCompileStatus(ctx context.Context, req *biz.UpdateCodeCompileStatusReq) (err error) {
	const sqlStr = "update judges set compile_status = ?, compile_log = ? where judge_id = ?"
	_, err = c.db.ExecContext(ctx, sqlStr,
		req.CompileStatus,
		req.CompileLog,
		req.CodeId,
	)
	return
}

func (c *codeProcessingImpl) UpdateCodeJudgeStatus(ctx context.Context, req *biz.UpdateCodeJudgeStatusReq) (err error) {
	const sqlStr = "update judges set compile_status = ?, judge_status = ?, cpu_time_used = ?, memory_used = ? where judge_id = ?"
	_, err = c.db.ExecContext(ctx, sqlStr,
		req.CompileStatus,
		req.JudgeStatus,
		req.CpuTimeUsed,
		req.MemoryUsed,
		req.CodeId,
	)
	return
}
