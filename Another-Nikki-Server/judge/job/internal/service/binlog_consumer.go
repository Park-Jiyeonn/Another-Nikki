package service

import (
	codeService "Another-Nikki/interact_hub/service/api"
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
	JudgeId       string `json:"judge_id"`
	CreatedTime   string `json:"created_time"`
	UserId        string `json:"user_id"`
	UserName      string `json:"user_name"`
	ProblemId     string `json:"problem_id"`
	ProblemName   string `json:"problem_name"`
	Language      string `json:"language"`
	Code          string `json:"code"`
	CompileLog    string `json:"compile_log"`
	CompileStatus string `json:"compile_status"`
	JudgeStatus   string `json:"judge_status"`
	CpuTimeUsed   string `json:"cpu_time_used"`
	MemoryUsed    string `json:"memory_used"`
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
	if msg == nil || len(msg.NewData) == 0 {
		log.Error(ctx, "binlog message was nil")
		return
	}
	if len(msg.NewData) != 1 {
		log.Error(ctx, "binlog message was not 1")
		return
	}
	return s.judge(ctx, msg.NewData[0])
}

func (s *JudgeBinlogConsumer) judge(ctx context.Context, data *Judge) (err error) {
	judgeId, err := strconv.ParseInt(data.JudgeId, 10, 64)
	if err != nil {
		log.Error(ctx, "judge err: %v", err)
		return
	}

	_, err = s.codeProcessingClient.UpdateCodeCompileStatus(ctx, &codeService.UpdateCodeCompileStatusReq{
		CodeId:        judgeId,
		CompileStatus: "Compiling",
		CompileLog:    "",
	})
	if err != nil {
		log.Error(ctx, "modify compile status err: %v", err)
		return
	}

	var judgeResp *judge.JudgeResp
	judgeResp, err = s.judgeClient.Judge(ctx, &judge.JudgeReq{
		Code:        data.Code,
		Language:    judge.Language(judge.Language_value[data.Language]),
		ProblemName: data.ProblemId + "." + data.ProblemName,
	})
	if err != nil {
		log.Error(ctx, "judge err: %v", err)
		_, _ = s.codeProcessingClient.UpdateCodeCompileStatus(ctx, &codeService.UpdateCodeCompileStatusReq{
			CodeId:        judgeId,
			CompileStatus: "未知错误",
			CompileLog:    err.Error(),
		})
		return
	}
	if judgeResp.IsCompileError {
		_, _ = s.codeProcessingClient.UpdateCodeCompileStatus(ctx, &codeService.UpdateCodeCompileStatusReq{
			CodeId:        judgeId,
			CompileStatus: judgeResp.CompileState,
			CompileLog:    judgeResp.CompileLog,
		})
		return
	}

	_, err = s.codeProcessingClient.UpdateCodeJudgeStatus(ctx, &codeService.UpdateCodeJudgeStatusReq{
		CodeId:        judgeId,
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
