//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"Another-Nikki/interact_hub/service/internal/client"
	"Another-Nikki/interact_hub/service/internal/conf"
	"Another-Nikki/interact_hub/service/internal/data_sqlite"
	"Another-Nikki/interact_hub/service/internal/server"
	"Another-Nikki/interact_hub/service/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger, *conf.Avatars, *conf.ClientTimeout) (*kratos.App, func(), error) {
	panic(wire.Build(service.ProviderSet, server.ProviderSet, client.GrpcProviderSet, data_sqlite.ProviderSet, newApp))
}
