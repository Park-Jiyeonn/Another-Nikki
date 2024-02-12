package server

import (
	"Another-Nikki/interact_hub/service/api"
	"Another-Nikki/interact_hub/service/internal/conf"
	"Another-Nikki/interact_hub/service/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func NewGRPCServer(c *conf.Server,
	problemService *service.ProblemService,
	articleService *service.ArticleService,
	commentService *service.CommentService,
	userService *service.UserService,
	logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			tracing.Server(), //设置trace，传入 trace provider
			logging.Server(logger),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	api.RegisterProblemServer(srv, problemService)
	api.RegisterArticleServer(srv, articleService)
	api.RegisterCommentServer(srv, commentService)
	api.RegisterUserServer(srv, userService)
	return srv
}
