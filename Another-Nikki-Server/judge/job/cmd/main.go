package main

import (
	zapLog "Another-Nikki/pkg/log"
	"Another-Nikki/pkg/trace"
	"flag"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/tx7do/kratos-transport/transport/kafka"
	"os"

	"Another-Nikki/judge/job/internal/conf"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	_ "go.uber.org/automaxprocs"
)

var (
	ServiceName = "Another-Nikki.Judge-Job"
	// flagconf is the config flag.
	flagconf string

	Env, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, job *kafka.Server, r registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.Name(ServiceName),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			job,
		),
		kratos.Registrar(r),
	)
}

func main() {
	flag.Parse()
	trace.Init(Env, ServiceName)
	logger := zapLog.Init(Env, ServiceName)
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
