package biz

import (
	"golang.org/x/net/context"
	"time"
)

type Problems struct {
	ProblemId    int64
	ProblemTitle string
	CreateTime   string
}

type ProblemRepo interface {
	PostProblem(ctx context.Context, req *PostProblemReq) (resp *PostProblemResp, err error)
	UpdateProblem(ctx context.Context, req *UpdateProblemReq) (err error)
	GetProblemById(ctx context.Context, req *GetProblemByIdReq) (*GetProblemByIdResp, error)
	GetProblemByPage(ctx context.Context, req *GetProblemByPageReq) (*GetProblemByPageResp, error)
}

type PostProblemReq struct {
	ProblemTitle       string `db:"problem_title"`
	ProblemDescription string `db:"problem_description"`
	ProblemContent     string `db:"problem_content"`
}
type PostProblemResp struct {
	ProblemId int64 `db:"problem_id"`
}

type UpdateProblemReq struct {
	ProblemId          int64  `db:"problem_id"`
	ProblemTitle       string `db:"problem_title"`
	ProblemDescription string `db:"problem_description"`
	ProblemContent     string `db:"problem_content"`
}

type GetProblemByIdReq struct {
	ProblemId int64 `db:"problem_id"`
}

type GetProblemByIdResp struct {
	ProblemTitle       string    `db:"problem_title"`
	ProblemDescription string    `db:"problem_description"`
	ProblemContent     string    `db:"problem_content"`
	CreatedTime        time.Time `db:"created_time"`
}

type GetProblemByPageReq struct {
	PageNum  int64 `db:"page_num"`
	PageSize int64 `db:"page_size"`
}

type ProblemPageDetail struct {
	ProblemId    int64     `db:"problem_id"`
	ProblemTitle string    `db:"problem_title"`
	CreatedTime  time.Time `db:"created_time"`
}

type GetProblemByPageResp struct {
	Problems []*ProblemPageDetail `db:"problems"`
}
