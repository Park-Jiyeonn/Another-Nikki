package service

import (
	"github.com/tx7do/kratos-transport/broker"
	"golang.org/x/net/context"
)

type JudgeBinlogConsumer struct {
}

type Judge struct {
	ID            string
	UserID        string `json:"user_id"`
	UserName      string `json:"user_name"`
	ProblemID     string `json:"problem_id"`
	ProblemName   string `json:"problem_name"`
	Language      string `json:"language"`
	Code          string `json:"code"`
	CompileStatus string `json:"compile_status"`
	CompileLog    string `json:"compile_log"`
	JudgeStatus   string `json:"judge_status"`
}

type Value struct {
	NewData  []*Judge `json:"data"`
	OldData  []*Judge `json:"old"`
	Type     string
	Table    string
	Database string
}

func NewJudgeBinlogConsumer() *JudgeBinlogConsumer {
	return &JudgeBinlogConsumer{}
}

func (s *JudgeBinlogConsumer) Handle(ctx context.Context, _ string, _ broker.Headers, msg *Value) (err error) {

	return
}
