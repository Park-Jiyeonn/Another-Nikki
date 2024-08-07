package server

import (
	"Another-Nikki/interact_hub/service/api"
	"Another-Nikki/interact_hub/service/internal/conf"
	"Another-Nikki/interact_hub/service/internal/service"
	pkg_jwt "Another-Nikki/pkg/jwt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/goccy/go-json"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/handlers"
	"golang.org/x/net/context"
	"strings"
)

func NewHTTPServer(c *conf.Server, logger log.Logger,
	problem *service.ProblemService,
	processingService *service.CodeProcessingService,
	userService *service.UserService,
	articleService *service.ArticleService,
	commentService *service.CommentService,
	logService *service.LogsService,
) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			getip(),
			tracing.Server(),
			// 日志相关的 api 不要记录日志
			selector.Server(
				logging.Server(logger),
			).Match(NotMatchLog()).Build(),
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
	api.RegisterCommentHTTPServer(srv, commentService)
	api.RegisterLogsHTTPServer(srv, logService)
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

type set struct {
	mp map[string]struct{}
}

func (s *set) insert(x ...string) {
	for i := 0; i < len(x); i++ {
		s.mp[x[i]] = struct{}{}
	}
}
func (s *set) count(x string) bool {
	_, ok := s.mp[x]
	return ok
}
func newSet() *set {
	return &set{mp: make(map[string]struct{})}
}

// NewWhiteListMatcher 白名单不需要token验证的接口
func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := newSet()

	// problem
	whiteList.insert("/service.problem.api.Problem/GetProblemById",
		"/service.problem.api.Problem/GetProblemByPage",
	)
	// user
	whiteList.insert("/service.problem.api.User/Register",
		"/service.problem.api.User/Login",
		"/service.problem.api.User/CreateTouristAccount",
		"/service.problem.api.User/GetUserById",
		"/service.problem.api.User/GetUserCommitRecordByPage",
		"/service.problem.api.User/GetUserSumCommit",
		"/service.problem.api.User/GetUserWrongProblem",
	)
	// article
	whiteList.insert("/service.problem.api.Article/GetArticleById",
		"/service.problem.api.Article/GetArticleByPage",
		"/service.problem.api.Comment/GetCommentsByArticleId",
		"/service.problem.api.Comment/GetLastSevenComment",
		"/service.problem.api.Comment/GetRandomComment",
	)

	return func(ctx context.Context, operation string) bool {
		if whiteList.count(operation) {
			return false
		}
		return true
	}
}

func getip() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				// 断言成HTTP的Transport可以拿到特殊信息
				if ht, ok := tr.(*http.Transport); ok {
					//fmt.Println(ht.Request())
					//fmt.Println(ht.Request().Header)
					//fmt.Printf("%v\n", ht.Request().RemoteAddr)
					//fmt.Println(ht.Request().Header.Get("Sec-Ch-Ua-Platform"))
					//fmt.Println(ht.Request().RequestURI)
					xForwardedFor := ht.Request().Header.Get("X-Forwarded-For")
					xForwardedFor = strings.ReplaceAll(xForwardedFor, "[", "")
					xForwardedFor = strings.ReplaceAll(xForwardedFor, "]", "")
					xForwardedFor = strings.ReplaceAll(xForwardedFor, "\"", "")
					ctx = context.WithValue(ctx, "ip", xForwardedFor)
					platform := ht.Request().Header.Get("Sec-Ch-Ua-Platform")
					if platform == "" {
						platform = ht.Request().Header.Get("User-Agent")
						n := len(platform)
						for i := 0; i < n; i++ {
							if platform[i] != '(' {
								continue
							}
							j := i
							for j+1 < n && platform[j+1] != ')' {
								j++
							}
							platform = platform[i+1 : j+1]
							break
						}
					}
					ctx = context.WithValue(ctx, "platform", platform)
					ctx = context.WithValue(ctx, "url", ht.Request().RequestURI)
				}
			}
			return handler(ctx, req)
		}
	}
}

func NotMatchLog() selector.MatchFunc {
	return func(ctx context.Context, operation string) bool {
		return !(strings.Contains(operation, "Log") || strings.Contains(operation, "log"))
	}
}
