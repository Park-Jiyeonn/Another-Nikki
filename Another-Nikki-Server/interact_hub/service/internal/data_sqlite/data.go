package data_sqlite

import (
	"Another-Nikki/interact_hub/service/internal/conf"
	"Another-Nikki/pkg/log"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/net/context"
)

var ProviderSet = wire.NewSet(NewData, NewSqliteDB,
	NewProblemRepo,
	NewArticleRepo,
	NewCommentImpl,
	NewUserImpl,
	NewCodeProcessingImpl,
	NewLogsRepoImpl,
)

type Data struct {
	GlobalDB *sqlx.DB
}

func NewData(db *sqlx.DB) (*Data, func(), error) {
	cleanup := func() {
		if err := db.Close(); err != nil {
			log.Info(context.Background(), "close DB client err: %v", err)
		}
	}
	return &Data{
		GlobalDB: db,
	}, cleanup, nil
}

func NewSqliteDB(c *conf.Data) *sqlx.DB {
	return sqlx.MustConnect(c.Database.Driver, c.Database.Source)
}
