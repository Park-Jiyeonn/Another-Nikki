//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"Another-Nikki/interact_hub/service/internal/conf"
	"Another-Nikki/interact_hub/service/internal/data"
	"Another-Nikki/interact_hub/service/internal/server"
	"Another-Nikki/interact_hub/service/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger, *conf.Avatars) (*kratos.App, func(), error) {
	panic(wire.Build(service.ProviderSet, server.ProviderSet, data.ProviderSet, newApp))
}
