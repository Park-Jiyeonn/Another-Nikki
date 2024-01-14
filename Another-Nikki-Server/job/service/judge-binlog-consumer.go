package service

import (
	"Another-Nikki/service/judge/api"
	"context"
	"encoding/json"
	"fmt"
	"log"
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
		log.Println(err)
	}
	switch val.Type {
	case "INSERT":
		fmt.Println("start to judge...")
		resp, err := s.globalGrpc.JudgeClient.Judge(context.Background(), &api.JudgeReq{
			Code:        val.NewData[0].Code,
			Language:    api.Language(api.Language_value[val.NewData[0].Language]),
			ProblemName: val.NewData[0].ProblemName,
		})
		fmt.Println(resp, err)
	case "UPDATE":
		return
	}
}
