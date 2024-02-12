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
	PostProblem(ctx context.Context, req *PostProblemReq) (err error)
	GetProblemById(ctx context.Context, req *GetProblemByIdReq) (*GetProblemByIdResp, error)
	GetProblemByPage(ctx context.Context, req *GetProblemByPageReq) (*GetProblemByPageResp, error)
}

type PostProblemReq struct {
	ProblemTitle       string `json:"problem_title"`
	ProblemDescription string `json:"problem_description"`
	ProblemContent     string `json:"problem_content"`
}

type GetProblemByIdReq struct {
	ProblemId int64 `json:"problem_id"`
}

type GetProblemByIdResp struct {
	ProblemTitle       string    `json:"problem_title"`
	ProblemDescription string    `json:"problem_description"`
	ProblemContent     string    `json:"problem_content"`
	CreateTime         time.Time `json:"create_time"`
}

type GetProblemByPageReq struct {
	PageNum  int64 `json:"page_num"`
	PageSize int64 `json:"page_size"`
}

type ProblemPageDetail struct {
	ProblemId    int64     `json:"problem_id"`
	ProblemTitle string    `json:"problem_title"`
	CreateTime   time.Time `json:"create_time"`
}

type GetProblemByPageResp struct {
	Problems []*ProblemPageDetail `json:"problems"`
}
