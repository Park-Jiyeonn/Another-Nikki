package data_sqlite

import (
	"Another-Nikki/interact_hub/service/internal/biz"
	"Another-Nikki/pkg/log"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
)

type codeProcessingImpl struct {
	db *sqlx.DB
}

func InitJudgesTable(db *sqlx.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS judges (
   judge_id INTEGER PRIMARY KEY AUTOINCREMENT,
   created_time DATETIME DEFAULT CURRENT_TIMESTAMP,
   updated_time DATETIME DEFAULT CURRENT_TIMESTAMP,
   user_id INTEGER NOT NULL DEFAULT 0,
   user_name TEXT NOT NULL DEFAULT '',
   problem_id INTEGER NOT NULL DEFAULT 0,
   problem_name TEXT NOT NULL DEFAULT '',
   language TEXT NOT NULL DEFAULT '',
   code TEXT NOT NULL,
   compile_status TEXT NOT NULL DEFAULT 'in queue',
   compile_log TEXT,
   judge_status TEXT NOT NULL DEFAULT '-',
   cpu_time_used TEXT NOT NULL DEFAULT '-',
   memory_used TEXT NOT NULL DEFAULT '-'
   );`
	_, err := db.Exec(schema)
	return err
}

func NewCodeProcessingImpl(data *Data) biz.CodeDataRepo {
	if err := InitJudgesTable(data.GlobalDB); err != nil {
		panic(err)
	}
	return &codeProcessingImpl{
		db: data.GlobalDB,
	}
}

func (c *codeProcessingImpl) CreateCode(ctx context.Context, req *biz.CreateCodeReq) (resp *biz.CreateCodeResp, err error) {
	resp = new(biz.CreateCodeResp)
	const sqlStr = `INSERT INTO judges 
		(user_id, user_name, problem_id, problem_name, language, code, updated_time) 
		VALUES (?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP)`
	row, err := c.db.ExecContext(ctx, sqlStr,
		req.UserId,
		req.UserName,
		req.ProblemId,
		req.ProblemName,
		req.Language,
		req.Code,
	)
	if err != nil {
		log.Error(ctx, "create judge code err: %v", err)
		return
	}
	resp.LastInsertId, err = row.LastInsertId()
	if err != nil {
		log.Error(ctx, "get LastInsertId err: %v", err)
		return
	}
	return
}

func (c *codeProcessingImpl) UpdateCodeCompileStatus(ctx context.Context, req *biz.UpdateCodeCompileStatusReq) (err error) {
	const sqlStr = `UPDATE judges 
		SET compile_status = ?, compile_log = ?, updated_time = CURRENT_TIMESTAMP 
		WHERE judge_id = ?`
	_, err = c.db.ExecContext(ctx, sqlStr,
		req.CompileStatus,
		req.CompileLog,
		req.CodeId,
	)
	return
}

func (c *codeProcessingImpl) UpdateCodeJudgeStatus(ctx context.Context, req *biz.UpdateCodeJudgeStatusReq) (err error) {
	const sqlStr = `UPDATE judges 
		SET compile_status = ?, judge_status = ?, cpu_time_used = ?, memory_used = ?, updated_time = CURRENT_TIMESTAMP 
		WHERE judge_id = ?`
	_, err = c.db.ExecContext(ctx, sqlStr,
		req.CompileStatus,
		req.JudgeStatus,
		req.CpuTimeUsed,
		req.MemoryUsed,
		req.CodeId,
	)
	return
}

func (c *codeProcessingImpl) GetCommitByJudgeId(ctx context.Context, req *biz.GetCommitByJudgeIdReq) (resp *biz.GetCommitByJudgeIdResp, err error) {
	const sqlStr = `SELECT problem_name, compile_status, judge_status, cpu_time_used, memory_used, language, 
		created_time, problem_id, code, user_name, user_id, compile_log 
		FROM judges 
		WHERE judge_id = ?`
	resp = new(biz.GetCommitByJudgeIdResp)
	err = c.db.GetContext(ctx, resp, sqlStr, req.JudgeId)
	return
}
