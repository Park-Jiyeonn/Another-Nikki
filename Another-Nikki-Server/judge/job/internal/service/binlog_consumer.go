package service

import (
	codeService "Another-Nikki/code_processing/service/api"
	"Another-Nikki/judge/job/internal/data"
	judge "Another-Nikki/judge/service/api"
	"Another-Nikki/pkg/log"
	"github.com/tx7do/kratos-transport/broker"
	"golang.org/x/net/context"
	"strconv"
)

type JudgeBinlogConsumer struct {
	judgeClient          judge.JudgeClient
	codeProcessingClient codeService.CodeProcessingClient
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
		judgeClient:          client.JudgeClient,
		codeProcessingClient: client.CodeProcessingClient,
	}
}

func (s *JudgeBinlogConsumer) Handle(ctx context.Context, _ string, _ broker.Headers, msg *Value) (err error) {
	//fmt.Println(msg)
	if msg.Type != "INSERT" {
		return
	}
	codeId, err := strconv.ParseInt(msg.NewData[0].ID, 10, 64)
	if err != nil {
		log.Error(ctx, "judge err: %v", err)
		return
	}

	_, err = s.codeProcessingClient.UpdateCodeCompileStatus(ctx, &codeService.UpdateCodeCompileStatusReq{
		CodeId:     codeId,
		Status:     "Compiling...",
		CompileLog: "",
	})
	if err != nil {
		log.Error(ctx, "modify compile status err: %v", err)
		return
	}

	var judgeResp *judge.JudgeResp
	judgeResp, err = s.judgeClient.Judge(ctx, &judge.JudgeReq{
		Code:        msg.NewData[0].Code,
		Language:    judge.Language(judge.Language_value[msg.NewData[0].Language]),
		ProblemName: msg.NewData[0].ProblemName,
	})
	if err != nil {
		log.Error(ctx, "judge err: %v", err)
	}
	if judgeResp.IsCompileError {
		_, _ = s.codeProcessingClient.UpdateCodeCompileStatus(ctx, &codeService.UpdateCodeCompileStatusReq{
			CodeId:     codeId,
			Status:     judgeResp.CompileState,
			CompileLog: judgeResp.CompileLog,
		})
		return
	}

	_, err = s.codeProcessingClient.UpdateCodeJudgeStatus(ctx, &codeService.UpdateCodeJudgeStatusReq{
		CodeId:        codeId,
		CompileStatus: judgeResp.CompileState,
		JudgeStatus:   judgeResp.JudgeResult,
		CpuTimeUsed:   judgeResp.CpuTimeUsed,
		MemoryUsed:    judgeResp.MemoryUsed,
	})
	if err != nil {
		log.Error(ctx, "modify judge resp err: %v", err)
		return
	}
	return
}
