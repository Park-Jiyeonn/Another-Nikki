//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"Another-Nikki/judge/service/internal/conf"
	"Another-Nikki/judge/service/internal/data"
	"Another-Nikki/judge/service/internal/server"
	"Another-Nikki/judge/service/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, service.ProviderSet, data.ProviderSet, newApp))
}
