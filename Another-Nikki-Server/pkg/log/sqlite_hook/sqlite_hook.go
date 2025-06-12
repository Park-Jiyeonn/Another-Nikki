package sqlite_hook

import (
	"github.com/goccy/go-json"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type SqliteHook struct {
	db *sqlx.DB
}

type LogEntry struct {
	Level       string `json:"level" db:"level"`
	Ts          string `json:"ts" db:"ts"`
	ServiceName string `json:"service.name" db:"service_name"`
	TraceID     string `json:"trace.id" db:"trace_id"`
	IP          string `json:"ip" db:"ip"`
	Platform    string `json:"platform" db:"platform"`
	URL         string `json:"url" db:"url"`
	Msg         string `json:"msg" db:"msg"`
	Args        string `json:"args" db:"args"`
	Stack       string `json:"stack" db:"stack"`
	Code        string `json:"code" db:"code"`
}

func NewSqliteHook() *SqliteHook {
	db := newSqlite()
	createTableIfNotExists(db)
	return &SqliteHook{db: db}
}

func newSqlite() *sqlx.DB {
	db := sqlx.MustConnect("sqlite3", "../../interact_hub/service/interact.db")
	return db
}

func createTableIfNotExists(db *sqlx.DB) {
	schema := `
	CREATE TABLE IF NOT EXISTS logs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		level TEXT,
		ts TEXT,
		service_name TEXT,
		trace_id TEXT,
		ip TEXT,
		platform TEXT,
		url TEXT,
		msg TEXT,
		args TEXT,
		stack TEXT,
		code TEXT
	);`
	db.MustExec(schema)
}

func (hook *SqliteHook) Write(p []byte) (n int, err error) {
	n = len(p)
	var logEntry LogEntry
	err = json.Unmarshal(p, &logEntry)
	if err != nil {
		return
	}

	// 解析时间格式
	var t time.Time
	t, err = time.Parse("2006-01-02T15:04:05.000-0700", logEntry.Ts)
	if err == nil {
		logEntry.Ts = t.Format("2006-01-02 15:04:05")
	}

	query := `
	INSERT INTO logs (level, ts, service_name, trace_id, ip, platform, url, msg, args, stack, code)
	VALUES (:level, :ts, :service_name, :trace_id, :ip, :platform, :url, :msg, :args, :stack, :code)
	`

	_, err = hook.db.NamedExec(query, logEntry)
	return
}
