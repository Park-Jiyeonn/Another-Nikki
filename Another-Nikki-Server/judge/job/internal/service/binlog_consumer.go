package service

import (
	"Another-Nikki/judge/job/internal/data"
	judge "Another-Nikki/judge/service/api"
	"Another-Nikki/pkg/log"
	"fmt"
	"github.com/tx7do/kratos-transport/broker"
	"golang.org/x/net/context"
)

type JudgeBinlogConsumer struct {
	judgeClient judge.JudgeClient
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

func NewJudgeBinlogConsumer(client *data.GlobalGrpcClient) *JudgeBinlogConsumer {
	return &JudgeBinlogConsumer{
		judgeClient: client.JudgeClient,
	}
}

func (s *JudgeBinlogConsumer) Handle(ctx context.Context, _ string, _ broker.Headers, msg *Value) (err error) {
	fmt.Println(msg)
	judgeResp, err := s.judgeClient.Judge(ctx, &judge.JudgeReq{
		Code:        msg.NewData[0].Code,
		Language:    judge.Language(judge.Language_value[msg.NewData[0].Language]),
		ProblemName: "1.A+B",
	})
	log.Info(ctx, "%+v, %+v", judgeResp, err)
	return
}
