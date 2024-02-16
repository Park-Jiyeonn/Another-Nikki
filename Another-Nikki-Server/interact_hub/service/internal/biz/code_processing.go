package biz

import (
	"golang.org/x/net/context"
)

type CodeDataRepo interface {
	CreateCode(ctx context.Context, req *CreateCodeReq) (err error)
	UpdateCodeCompileStatus(ctx context.Context, req *UpdateCodeCompileStatusReq) (err error)
	UpdateCodeJudgeStatus(ctx context.Context, req *UpdateCodeJudgeStatusReq) (err error)
}

type CreateCodeReq struct {
	UserId      int64  `db:"user_id"`
	UserName    string `db:"user_name"`
	ProblemId   int64  `db:"problem_id"`
	ProblemName string `db:"problem_name"`
	Language    string `db:"language"`
	Code        string `db:"code"`
}

type UpdateCodeCompileStatusReq struct {
	CodeId        int64  `db:"code_id"`
	CompileStatus string `db:"compile_status"`
	CompileLog    string `db:"compile_log"`
}

type UpdateCodeJudgeStatusReq struct {
	CodeId        int64  `db:"code_id"`
	CompileStatus string `db:"compile_status"`
	JudgeStatus   string `db:"judge_status"`
	CpuTimeUsed   string `db:"cpu_time_used"`
	MemoryUsed    string `db:"memory_used"`
}