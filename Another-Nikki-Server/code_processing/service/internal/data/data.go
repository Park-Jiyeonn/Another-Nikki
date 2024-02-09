package data

import (
	"Another-Nikki/code_processing/service/internal/conf"
	"Another-Nikki/pkg/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
)

var ProviderSet = wire.NewSet(NewData, NewGlobalGrpcClient, NewRegistry, NewDiscovery, NewMySql, NewCodeProcessingRepo)

type Data struct {
	DB *sqlx.DB
}

func NewData(c *conf.Data, db *sqlx.DB) (*Data, func(), error) {
	cleanup := func() {
		if err := db.Close(); err != nil {
			log.Info(context.Background(), "close DB client err: %v", err)
		}
	}
	return &Data{
		DB: db,
	}, cleanup, nil
}

func NewMySql(c *conf.Data) *sqlx.DB {
	db, err := sqlx.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		panic(err)
	}
	return db
}
