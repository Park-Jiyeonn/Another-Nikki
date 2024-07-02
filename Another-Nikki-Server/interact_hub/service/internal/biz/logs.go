package biz

import (
	"golang.org/x/net/context"
	"time"
)

type LogsRepo interface {
	GetLogs(ctx context.Context, req *GetLogsReq) (*GetLogsResp, error)
}

type GetLogsReq struct {
	PageNum  int64 `db:"page_num"`
	PageSize int64 `db:"page_size"`
}

type LogDetail struct {
	LogId       int64     `json:"log_id" db:"log_id"`
	Level       string    `json:"level" db:"level"`
	Ts          time.Time `json:"ts" db:"ts"`
	ServiceName string    `json:"service.name" db:"service_name"`
	TraceID     string    `json:"trace.id" db:"trace_id"`
	IP          string    `json:"ip" db:"ip"`
	Platform    string    `json:"platform" db:"platform"`
	URL         string    `json:"url" db:"url"`
	Msg         string    `json:"msg" db:"msg"`
	Args        string    `json:"args" db:"args"`
	Stack       string    `json:"stack" db:"stack"`
	Code        string    `json:"code" db:"code"`
}

type GetLogsResp struct {
	SumLog int64        `json:"sum_log" db:"sum_log"`
	Logs   []*LogDetail `json:"logs" db:"logs"`
}
