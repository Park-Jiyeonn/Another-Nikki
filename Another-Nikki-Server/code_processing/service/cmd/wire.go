//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"Another-Nikki/code_processing/service/internal/biz"
	"Another-Nikki/code_processing/service/internal/conf"
	"Another-Nikki/code_processing/service/internal/data"
	"Another-Nikki/code_processing/service/internal/server"
	"Another-Nikki/code_processing/service/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(service.ProviderSet, server.ProviderSet, data.ProviderSet, biz.ProviderSet, newApp))
}
