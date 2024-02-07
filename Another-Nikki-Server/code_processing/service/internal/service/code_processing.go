package service

import (
	pb "Another-Nikki/code_processing/service/api"
	"Another-Nikki/code_processing/service/internal/data"
	judge "Another-Nikki/judge/service/api"
	"Another-Nikki/pkg/log"
	"context"
)

type CodeProcessingService struct {
	pb.UnimplementedCodeProcessingServer

	judgeClient judge.JudgeClient
}

func NewCodeProcessingService(data *data.Data) *CodeProcessingService {
	return &CodeProcessingService{
		judgeClient: data.GlobalGrpcClient.JudgeClient,
	}
}

func (s *CodeProcessingService) SubmitCode(ctx context.Context, req *pb.SubmitCodeReq) (*pb.SubmitCodeResp, error) {
	resp, err := s.judgeClient.Judge(ctx, &judge.JudgeReq{
		Code:        "1234",
		Language:    judge.Language_Java,
		ProblemName: "123",
	})
	//log.Context(ctx).Infow("resp = ", resp, "err = ", err)
	log.Info(ctx, "resp", resp, "err", err)
	//log.Context(ctx).Log(log.LevelError, "resp = ", resp, "err = ", err)
	//_, _ = s.judgeClient.OnlineRun(ctx, &judge.OnlineRunReq{
	//	Code:     "1234",
	//	Language: judge.Language_Python,
	//	Input:    "",
	//})
	//fmt.Println(ret, err)
	return &pb.SubmitCodeResp{}, nil
}
