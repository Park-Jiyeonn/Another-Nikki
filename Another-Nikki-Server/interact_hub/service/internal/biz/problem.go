package biz

import "golang.org/x/net/context"

type Problems struct {
	ProblemId    int64
	ProblemTitle string
	CreateTime   string
}
type ProblemRepo interface {
	CreateProblem(ctx context.Context, problemId int64, title, description, content string) (err error)
	GetProblemById(ctx context.Context, problemId int64) (problemTitle, problemDescription, problemContent, createTime string, err error)
	GetProblemByPage(ctx context.Context, pageNum, pageSize int64) (problems []*Problems, err error)
}
