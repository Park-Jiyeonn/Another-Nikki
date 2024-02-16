package server

import (
	"Another-Nikki/interact_hub/service/api"
	"Another-Nikki/interact_hub/service/internal/conf"
	"Another-Nikki/interact_hub/service/internal/service"
	pkg_jwt "Another-Nikki/pkg/jwt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/goccy/go-json"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/handlers"
	"golang.org/x/net/context"
)

func NewHTTPServer(c *conf.Server, logger log.Logger,
	problem *service.ProblemService,
	processingService *service.CodeProcessingService,
	userService *service.UserService,
	articleService *service.ArticleService,
) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
			selector.Server(
				jwt.Server(func(token *jwt2.Token) (interface{}, error) {
					return []byte(pkg_jwt.JwtSecret), nil
				}, jwt.WithSigningMethod(jwt2.SigningMethodHS256)),
			).Match(NewWhiteListMatcher()).Build(),
		),
		http.ResponseEncoder(responseEncoder),
		http.ErrorEncoder(ErrorEncoder),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}

	opts = append(opts, http.Filter(handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)))
	srv := http.NewServer(opts...)
	api.RegisterProblemHTTPServer(srv, problem)
	api.RegisterCodeProcessingHTTPServer(srv, processingService)
	api.RegisterUserHTTPServer(srv, userService)
	api.RegisterArticleHTTPServer(srv, articleService)
	return srv
}

func responseEncoder(w http.ResponseWriter, r *http.Request, data interface{}) error {
	type Response struct {
		Code int         `json:"code"`
		Data interface{} `json:"data"`
	}
	res := &Response{
		Code: 200,
		Data: data,
	}
	msRes, err := json.Marshal(res)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(msRes)
	return nil
}

type ErrResponse struct {
	Code    int32       `json:"code" form:"code"`
	Message string      `json:"message" form:"message"`
	Reason  string      `json:"reason" form:"reason"`
	Data    interface{} `json:"data" form:"data"`
}

func ErrorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	se := errors.FromError(err)
	reply := &ErrResponse{}
	reply.Code = se.Code
	reply.Data = nil
	reply.Message = se.Message
	reply.Reason = se.Reason

	body, err := json.Marshal(reply)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(body)
}

// NewWhiteListMatcher 白名单不需要token验证的接口
func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	whiteList["/service.problem.api.Problem/GetProblemById"] = struct{}{}
	whiteList["/service.problem.api.Problem/GetProblemByPage"] = struct{}{}
	whiteList["/service.problem.api.User/Register"] = struct{}{}
	whiteList["/service.problem.api.User/Login"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}
