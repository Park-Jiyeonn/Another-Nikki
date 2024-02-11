package server

import (
	"Another-Nikki/judge/job/internal/service"
	"github.com/tx7do/kratos-transport/transport/kafka"
	"golang.org/x/net/context"
)

func NewKafkaServer(svc *service.JudgeBinlogConsumer) *kafka.Server {
	ctx := context.Background()

	srv := kafka.NewServer(
		kafka.WithAddress([]string{"localhost:9092"}),
		kafka.WithGlobalTracerProvider(),
		kafka.WithGlobalPropagator(),
		kafka.WithCodec("json"),
	)

	registerKafkaSubscribers(ctx, srv, svc)

	return srv
}

func registerKafkaSubscribers(ctx context.Context, srv *kafka.Server, svc *service.JudgeBinlogConsumer) {
	err := kafka.RegisterSubscriber(srv, ctx,
		"hello", "consumer-group-id", false,
		svc.Handle,
	)
	if err != nil {
		panic(err)
	}
}
