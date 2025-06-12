package data_sqlite

import (
	"Another-Nikki/interact_hub/service/internal/biz"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
)

type logsImpl struct {
	db *sqlx.DB
}

func InitLogsTable(db *sqlx.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS logs (
		log_id INTEGER PRIMARY KEY AUTOINCREMENT,
		level TEXT,
		ts DATETIME,
		service_name TEXT,
		trace_id TEXT,
		ip TEXT,
		platform TEXT,
		url TEXT NOT NULL,
		msg TEXT NOT NULL,
		args TEXT NOT NULL,
		stack TEXT NOT NULL,
		code TEXT
	);`
	_, err := db.Exec(schema)
	return err
}

func NewLogsRepoImpl(data *Data) biz.LogsRepo {
	if err := InitLogsTable(data.GlobalDB); err != nil {
		panic(err)
	}
	return &logsImpl{
		db: data.GlobalDB,
	}
}

func (s *logsImpl) GetLogs(ctx context.Context, req *biz.GetLogsReq) (resp *biz.GetLogsResp, err error) {
	resp = new(biz.GetLogsResp)

	rows, err := s.db.QueryxContext(ctx,
		`SELECT log_id, level, ts, service_name, trace_id, ip, platform, url, msg, args, stack, code 
		 FROM logs 
		 ORDER BY log_id DESC 
		 LIMIT ? OFFSET ?`,
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

	// 获取总日志数量
	err = s.db.QueryRowContext(ctx, "SELECT COUNT(log_id) FROM logs").Scan(&resp.SumLog)
	return
}
