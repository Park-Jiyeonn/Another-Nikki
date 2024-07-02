package service

import (
	"Another-Nikki/interact_hub/service/internal/biz"
	"Another-Nikki/pkg/jwt"
	"context"
	"fmt"
	"time"

	pb "Another-Nikki/interact_hub/service/api"
)

type LogsService struct {
	pb.UnimplementedLogsServer
	dao biz.LogsRepo
}

func NewLogsService(dao biz.LogsRepo) *LogsService {
	return &LogsService{dao: dao}
}

func (s *LogsService) GetLogs(ctx context.Context, req *pb.GetLogsReq) (resp *pb.GetLogsResp, err error) {
	resp = new(pb.GetLogsResp)
	userId, _ := jwt.GetUserFromCtx(ctx)
	if userId > 2 {
		err = fmt.Errorf("访问错误，非管理员")
		return
	}
	var logs *biz.GetLogsResp
	logs, err = s.dao.GetLogs(ctx, &biz.GetLogsReq{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	})
	if err != nil {
		return
	}
	resp.Logs = make([]*pb.LogDetail, 0, len(logs.Logs))
	for _, val := range logs.Logs {
		resp.Logs = append(resp.Logs, &pb.LogDetail{
			Level:       val.Level,
			Ts:          val.Ts.Format(time.DateTime),
			ServiceName: val.ServiceName,
			TraceId:     val.TraceID,
			Ip:          val.IP,
			Platform:    val.Platform,
			Url:         val.URL,
			Msg:         val.Msg,
			Args:        val.Args,
			Stack:       val.Stack,
			Code:        val.Code,
			LogId:       val.LogId,
		})
	}
	resp.SumLog = logs.SumLog
	return
}
