package data

import (
	"Another-Nikki/interact_hub/service/internal/conf"
	"Another-Nikki/pkg/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
)

var ProviderSet = wire.NewSet(NewData, NewGlobalGrpcClient, NewRegistry, NewDiscovery, NewMySql,
	NewProblemRepo,
	NewArticleRepo,
	NewCommentImpl,
	NewUserImpl,
	NewCodeProcessingImpl,
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

func NewMySql(c *conf.Data) *sqlx.DB {
	return sqlx.MustConnect(c.Database.Driver, c.Database.Source)
}
