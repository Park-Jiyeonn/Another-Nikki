package log

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/goccy/go-json"
	"github.com/jmoiron/sqlx"
	"time"
)

type MySqlHook struct {
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

func NewMySqlHook() *MySqlHook {
	return &MySqlHook{db: newMySql()}
}

//CREATE TABLE logs (
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

//ALTER TABLE comments ADD COLUMN code VARCHAR(50);

func newMySql() *sqlx.DB {
	return sqlx.MustConnect("mysql", "jiyeon:1234@tcp(127.0.0.1:3307)/Another_Nikki?charset=utf8mb4&parseTime=True&loc=Local")
}
func (hook *MySqlHook) Write(p []byte) (n int, err error) {
	n = len(p)
	//fmt.Println(string(p))
	var logEntry LogEntry
	err = json.Unmarshal(p, &logEntry)
	var t time.Time
	t, err = time.Parse("2006-01-02T15:04:05.000-0700", logEntry.Ts)
	logEntry.Ts = t.Format("2006-01-02 15:04:05")
	query := `
    INSERT INTO logs (level, ts, service_name, trace_id, ip, platform, url, msg, args, stack, code)
    VALUES (:level, :ts, :service_name, :trace_id, :ip, :platform, :url, :msg, :args, :stack, :code)
    `
	_, err = hook.db.NamedExec(query, logEntry)
	return
}
