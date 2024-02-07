package service

import (
	"Another-Nikki/judge/service/api"
	"Another-Nikki/pkg/log"
	"context"
	"encoding/json"
	"fmt"
)

type Judge struct {
	ID            string
	UserID        string `json:"user_id"`
	UserName      string `json:"user_name"`
	ProblemID     string `json:"problem_id"`
	ProblemName   string `json:"problem_name"`
	Language      string `json:"language"`
	Code          string `json:"code"`
	CompileStatus string `json:"compile_status"`
	JudgeStatus   string `json:"judge_status"`
}

type Value struct {
	NewData  []*Judge `json:"data"`
	OldData  []*Judge `json:"old"`
	Type     string
	Table    string
	Database string
}

func (s *Service) MessageHandle(value []byte) {
	val := new(Value)
	err := json.Unmarshal(value, val)
	if err != nil {
		log.Error(context.Background(), "123")
	}
	switch val.Type {
	case "INSERT":
		resp, err := s.GlobalGrpc.JudgeClient.Judge(context.Background(), &api.JudgeReq{
			Code:        val.NewData[0].Code,
			Language:    api.Language(api.Language_value[val.NewData[0].Language]),
			ProblemName: val.NewData[0].ProblemName,
		})
		fmt.Println(resp, err)
		log.Info(context.Background(), "123")
	case "UPDATE":
		return
	}
}
