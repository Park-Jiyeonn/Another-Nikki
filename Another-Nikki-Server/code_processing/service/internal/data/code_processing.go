package data

import (
	"Another-Nikki/code_processing/service/api"
	"Another-Nikki/code_processing/service/internal/biz"
	"fmt"
	"golang.org/x/net/context"
	"time"
)

type CodeDataMysql struct {
	ID            int64 `gorm:"primarykey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	UserID        int64  `json:"user_id"`
	UserName      string `json:"user_name"`
	ProblemID     int64  `json:"problem_id"`
	ProblemName   string `json:"problem_name"`
	Language      string `json:"language"`
	Code          string `json:"code"`
	CompileStatus string `json:"compile_status"`
	JudgeStatus   string `json:"judge_status"`
}

type codeProcessingRepo struct {
	data *Data
}

func (c *codeProcessingRepo) CreateCode(ctx context.Context, req *api.SubmitCodeReq) (err error) {
	fmt.Println("yes, I am work")
	return
}

func NewCodeProcessingRepo(data *Data) biz.CodeDataRepo {
	return &codeProcessingRepo{
		data: data,
	}
}
