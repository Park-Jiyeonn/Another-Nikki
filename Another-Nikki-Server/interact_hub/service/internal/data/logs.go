package data

import (
	"Another-Nikki/interact_hub/service/internal/biz"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
)

type logsImpl struct {
	db *sqlx.DB
}

// CREATE TABLE logs (
//	log_id SERIAL PRIMARY KEY,
//	level VARCHAR(50),
//	ts TIMESTAMP,
//	service_name VARCHAR(100),
//	trace_id VARCHAR(100),
//	ip VARCHAR(45),
//	platform VARCHAR(50),
//	url TEXT NOT NULL,
//	msg TEXT NOT NULL,
//	args TEXT NOT NULL,
//	stack TEXT NOT NULL,
//	code VARCHAR(50)
//);

func NewLogsRepoImpl(data *Data) biz.LogsRepo {
	return &logsImpl{
		db: data.GlobalDB,
	}
}

func (s *logsImpl) GetLogs(ctx context.Context, req *biz.GetLogsReq) (resp *biz.GetLogsResp, err error) {
	resp = new(biz.GetLogsResp)
	rows, err := s.db.QueryxContext(ctx, "SELECT log_id, level, ts, service_name, trace_id, ip, platform, url, msg, args, stack, code FROM logs ORDER BY log_id DESC LIMIT ? OFFSET ?",
		req.PageSize, (req.PageNum-1)*req.PageSize)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var log biz.LogDetail
		err = rows.StructScan(&log)
		if err != nil {
			return nil, err
		}
		resp.Logs = append(resp.Logs, &log)
	}

	err = s.db.QueryRowContext(ctx, "SELECT COUNT(log_id) from logs").Scan(&resp.SumLog)
	return
}
