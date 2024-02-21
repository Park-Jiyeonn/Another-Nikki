// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.2
// - protoc             v4.24.4
// source: interact_hub.proto

package api

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationProblemGetProblemById = "/service.problem.api.Problem/GetProblemById"
const OperationProblemGetProblemByPage = "/service.problem.api.Problem/GetProblemByPage"
const OperationProblemPostProblem = "/service.problem.api.Problem/PostProblem"

type ProblemHTTPServer interface {
	GetProblemById(context.Context, *GetProblemByIdReq) (*GetProblemByIdResp, error)
	GetProblemByPage(context.Context, *GetProblemByPageReq) (*GetProblemByPageResp, error)
	PostProblem(context.Context, *PostProblemReq) (*PostProblemResp, error)
}

func RegisterProblemHTTPServer(s *http.Server, srv ProblemHTTPServer) {
	r := s.Route("/")
	r.POST("/api/problem/post", _Problem_PostProblem0_HTTP_Handler(srv))
	r.GET("/api/problem/{problem_id}", _Problem_GetProblemById0_HTTP_Handler(srv))
	r.GET("/api/problem/{page_size}/{page_num}", _Problem_GetProblemByPage0_HTTP_Handler(srv))
}

func _Problem_PostProblem0_HTTP_Handler(srv ProblemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PostProblemReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationProblemPostProblem)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PostProblem(ctx, req.(*PostProblemReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PostProblemResp)
		return ctx.Result(200, reply)
	}
}

func _Problem_GetProblemById0_HTTP_Handler(srv ProblemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetProblemByIdReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationProblemGetProblemById)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetProblemById(ctx, req.(*GetProblemByIdReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetProblemByIdResp)
		return ctx.Result(200, reply)
	}
}

func _Problem_GetProblemByPage0_HTTP_Handler(srv ProblemHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetProblemByPageReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationProblemGetProblemByPage)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetProblemByPage(ctx, req.(*GetProblemByPageReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetProblemByPageResp)
		return ctx.Result(200, reply)
	}
}

type ProblemHTTPClient interface {
	GetProblemById(ctx context.Context, req *GetProblemByIdReq, opts ...http.CallOption) (rsp *GetProblemByIdResp, err error)
	GetProblemByPage(ctx context.Context, req *GetProblemByPageReq, opts ...http.CallOption) (rsp *GetProblemByPageResp, err error)
	PostProblem(ctx context.Context, req *PostProblemReq, opts ...http.CallOption) (rsp *PostProblemResp, err error)
}

type ProblemHTTPClientImpl struct {
	cc *http.Client
}

func NewProblemHTTPClient(client *http.Client) ProblemHTTPClient {
	return &ProblemHTTPClientImpl{client}
}

func (c *ProblemHTTPClientImpl) GetProblemById(ctx context.Context, in *GetProblemByIdReq, opts ...http.CallOption) (*GetProblemByIdResp, error) {
	var out GetProblemByIdResp
	pattern := "/api/problem/{problem_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationProblemGetProblemById))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ProblemHTTPClientImpl) GetProblemByPage(ctx context.Context, in *GetProblemByPageReq, opts ...http.CallOption) (*GetProblemByPageResp, error) {
	var out GetProblemByPageResp
	pattern := "/api/problem/{page_size}/{page_num}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationProblemGetProblemByPage))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ProblemHTTPClientImpl) PostProblem(ctx context.Context, in *PostProblemReq, opts ...http.CallOption) (*PostProblemResp, error) {
	var out PostProblemResp
	pattern := "/api/problem/post"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationProblemPostProblem))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

const OperationArticleGetArticleById = "/service.problem.api.Article/GetArticleById"
const OperationArticleGetArticleByPage = "/service.problem.api.Article/GetArticleByPage"
const OperationArticlePostArticle = "/service.problem.api.Article/PostArticle"

type ArticleHTTPServer interface {
	GetArticleById(context.Context, *GetArticleByIdReq) (*GetArticleByIdResp, error)
	GetArticleByPage(context.Context, *GetArticleByPageReq) (*GetArticleByPageResp, error)
	PostArticle(context.Context, *PostArticleReq) (*PostArticleResp, error)
}

func RegisterArticleHTTPServer(s *http.Server, srv ArticleHTTPServer) {
	r := s.Route("/")
	r.POST("/api/article/post", _Article_PostArticle0_HTTP_Handler(srv))
	r.GET("/api/article/{article_id}", _Article_GetArticleById0_HTTP_Handler(srv))
	r.GET("/api/article/{page_num}/{page_size}", _Article_GetArticleByPage0_HTTP_Handler(srv))
}

func _Article_PostArticle0_HTTP_Handler(srv ArticleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PostArticleReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationArticlePostArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PostArticle(ctx, req.(*PostArticleReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PostArticleResp)
		return ctx.Result(200, reply)
	}
}

func _Article_GetArticleById0_HTTP_Handler(srv ArticleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetArticleByIdReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationArticleGetArticleById)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetArticleById(ctx, req.(*GetArticleByIdReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetArticleByIdResp)
		return ctx.Result(200, reply)
	}
}

func _Article_GetArticleByPage0_HTTP_Handler(srv ArticleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetArticleByPageReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationArticleGetArticleByPage)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetArticleByPage(ctx, req.(*GetArticleByPageReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetArticleByPageResp)
		return ctx.Result(200, reply)
	}
}

type ArticleHTTPClient interface {
	GetArticleById(ctx context.Context, req *GetArticleByIdReq, opts ...http.CallOption) (rsp *GetArticleByIdResp, err error)
	GetArticleByPage(ctx context.Context, req *GetArticleByPageReq, opts ...http.CallOption) (rsp *GetArticleByPageResp, err error)
	PostArticle(ctx context.Context, req *PostArticleReq, opts ...http.CallOption) (rsp *PostArticleResp, err error)
}

type ArticleHTTPClientImpl struct {
	cc *http.Client
}

func NewArticleHTTPClient(client *http.Client) ArticleHTTPClient {
	return &ArticleHTTPClientImpl{client}
}

func (c *ArticleHTTPClientImpl) GetArticleById(ctx context.Context, in *GetArticleByIdReq, opts ...http.CallOption) (*GetArticleByIdResp, error) {
	var out GetArticleByIdResp
	pattern := "/api/article/{article_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationArticleGetArticleById))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ArticleHTTPClientImpl) GetArticleByPage(ctx context.Context, in *GetArticleByPageReq, opts ...http.CallOption) (*GetArticleByPageResp, error) {
	var out GetArticleByPageResp
	pattern := "/api/article/{page_num}/{page_size}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationArticleGetArticleByPage))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ArticleHTTPClientImpl) PostArticle(ctx context.Context, in *PostArticleReq, opts ...http.CallOption) (*PostArticleResp, error) {
	var out PostArticleResp
	pattern := "/api/article/post"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationArticlePostArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

const OperationUserCreateTouristAccount = "/service.problem.api.User/CreateTouristAccount"
const OperationUserGetUserById = "/service.problem.api.User/GetUserById"
const OperationUserGetUserCommitRecordByPage = "/service.problem.api.User/GetUserCommitRecordByPage"
const OperationUserGetUserSumCommit = "/service.problem.api.User/GetUserSumCommit"
const OperationUserLogin = "/service.problem.api.User/Login"
const OperationUserRegister = "/service.problem.api.User/Register"
const OperationUserUpdateUser = "/service.problem.api.User/UpdateUser"

type UserHTTPServer interface {
	CreateTouristAccount(context.Context, *CreateTouristAccountReq) (*CreateTouristAccountResp, error)
	GetUserById(context.Context, *GetUserByIdReq) (*GetUserByIdResp, error)
	GetUserCommitRecordByPage(context.Context, *GetUserCommitRecordReq) (*GetUserCommitRecordResp, error)
	GetUserSumCommit(context.Context, *GetUserSumCommitReq) (*GetUserSumCommitResp, error)
	Login(context.Context, *LoginReq) (*LoginResp, error)
	Register(context.Context, *RegisterReq) (*RegisterResp, error)
	UpdateUser(context.Context, *UpdateUserReq) (*UpdateUserResp, error)
}

func RegisterUserHTTPServer(s *http.Server, srv UserHTTPServer) {
	r := s.Route("/")
	r.POST("/api/user/login", _User_Login0_HTTP_Handler(srv))
	r.POST("/api/user/register", _User_Register0_HTTP_Handler(srv))
	r.GET("/api/user/profile/{user_id}", _User_GetUserById0_HTTP_Handler(srv))
	r.GET("/api/user/profile/{user_id}/commit-record/{page_num}/{page_size}", _User_GetUserCommitRecordByPage0_HTTP_Handler(srv))
	r.GET("/api/user/profile/{user_id}/commit-record/sum", _User_GetUserSumCommit0_HTTP_Handler(srv))
	r.POST("/api/user/update", _User_UpdateUser0_HTTP_Handler(srv))
	r.POST("/api/user/create/tourist", _User_CreateTouristAccount0_HTTP_Handler(srv))
}

func _User_Login0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginResp)
		return ctx.Result(200, reply)
	}
}

func _User_Register0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserRegister)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RegisterResp)
		return ctx.Result(200, reply)
	}
}

func _User_GetUserById0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserByIdReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserGetUserById)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserById(ctx, req.(*GetUserByIdReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserByIdResp)
		return ctx.Result(200, reply)
	}
}

func _User_GetUserCommitRecordByPage0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserCommitRecordReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserGetUserCommitRecordByPage)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserCommitRecordByPage(ctx, req.(*GetUserCommitRecordReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserCommitRecordResp)
		return ctx.Result(200, reply)
	}
}

func _User_GetUserSumCommit0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserSumCommitReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserGetUserSumCommit)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserSumCommit(ctx, req.(*GetUserSumCommitReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserSumCommitResp)
		return ctx.Result(200, reply)
	}
}

func _User_UpdateUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateUserReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserUpdateUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUser(ctx, req.(*UpdateUserReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateUserResp)
		return ctx.Result(200, reply)
	}
}

func _User_CreateTouristAccount0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateTouristAccountReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserCreateTouristAccount)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateTouristAccount(ctx, req.(*CreateTouristAccountReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateTouristAccountResp)
		return ctx.Result(200, reply)
	}
}

type UserHTTPClient interface {
	CreateTouristAccount(ctx context.Context, req *CreateTouristAccountReq, opts ...http.CallOption) (rsp *CreateTouristAccountResp, err error)
	GetUserById(ctx context.Context, req *GetUserByIdReq, opts ...http.CallOption) (rsp *GetUserByIdResp, err error)
	GetUserCommitRecordByPage(ctx context.Context, req *GetUserCommitRecordReq, opts ...http.CallOption) (rsp *GetUserCommitRecordResp, err error)
	GetUserSumCommit(ctx context.Context, req *GetUserSumCommitReq, opts ...http.CallOption) (rsp *GetUserSumCommitResp, err error)
	Login(ctx context.Context, req *LoginReq, opts ...http.CallOption) (rsp *LoginResp, err error)
	Register(ctx context.Context, req *RegisterReq, opts ...http.CallOption) (rsp *RegisterResp, err error)
	UpdateUser(ctx context.Context, req *UpdateUserReq, opts ...http.CallOption) (rsp *UpdateUserResp, err error)
}

type UserHTTPClientImpl struct {
	cc *http.Client
}

func NewUserHTTPClient(client *http.Client) UserHTTPClient {
	return &UserHTTPClientImpl{client}
}

func (c *UserHTTPClientImpl) CreateTouristAccount(ctx context.Context, in *CreateTouristAccountReq, opts ...http.CallOption) (*CreateTouristAccountResp, error) {
	var out CreateTouristAccountResp
	pattern := "/api/user/create/tourist"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserCreateTouristAccount))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) GetUserById(ctx context.Context, in *GetUserByIdReq, opts ...http.CallOption) (*GetUserByIdResp, error) {
	var out GetUserByIdResp
	pattern := "/api/user/profile/{user_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserGetUserById))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) GetUserCommitRecordByPage(ctx context.Context, in *GetUserCommitRecordReq, opts ...http.CallOption) (*GetUserCommitRecordResp, error) {
	var out GetUserCommitRecordResp
	pattern := "/api/user/profile/{user_id}/commit-record/{page_num}/{page_size}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserGetUserCommitRecordByPage))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) GetUserSumCommit(ctx context.Context, in *GetUserSumCommitReq, opts ...http.CallOption) (*GetUserSumCommitResp, error) {
	var out GetUserSumCommitResp
	pattern := "/api/user/profile/{user_id}/commit-record/sum"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserGetUserSumCommit))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) Login(ctx context.Context, in *LoginReq, opts ...http.CallOption) (*LoginResp, error) {
	var out LoginResp
	pattern := "/api/user/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) Register(ctx context.Context, in *RegisterReq, opts ...http.CallOption) (*RegisterResp, error) {
	var out RegisterResp
	pattern := "/api/user/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserRegister))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...http.CallOption) (*UpdateUserResp, error) {
	var out UpdateUserResp
	pattern := "/api/user/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserUpdateUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

const OperationCommentGetCommentsByArticleId = "/service.problem.api.Comment/GetCommentsByArticleId"
const OperationCommentGetLastSevenComment = "/service.problem.api.Comment/GetLastSevenComment"
const OperationCommentGetRandomComment = "/service.problem.api.Comment/GetRandomComment"
const OperationCommentPostComment = "/service.problem.api.Comment/PostComment"

type CommentHTTPServer interface {
	GetCommentsByArticleId(context.Context, *GetCommentsByArticleIdReq) (*GetCommentsByArticleIdResp, error)
	GetLastSevenComment(context.Context, *GetLastSevenCommentReq) (*GetLastSevenCommentResp, error)
	GetRandomComment(context.Context, *GetRandomCommentReq) (*GetRandomCommentResp, error)
	PostComment(context.Context, *PostCommentReq) (*PostCommentResp, error)
}

func RegisterCommentHTTPServer(s *http.Server, srv CommentHTTPServer) {
	r := s.Route("/")
	r.POST("/api/comment/post", _Comment_PostComment0_HTTP_Handler(srv))
	r.GET("/api/comment/{article_id}", _Comment_GetCommentsByArticleId0_HTTP_Handler(srv))
	r.GET("/api/comment/last_seven/{article_id}/{num}", _Comment_GetLastSevenComment0_HTTP_Handler(srv))
	r.GET("/api/comment/random/{article_id}", _Comment_GetRandomComment0_HTTP_Handler(srv))
}

func _Comment_PostComment0_HTTP_Handler(srv CommentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PostCommentReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCommentPostComment)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PostComment(ctx, req.(*PostCommentReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PostCommentResp)
		return ctx.Result(200, reply)
	}
}

func _Comment_GetCommentsByArticleId0_HTTP_Handler(srv CommentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCommentsByArticleIdReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCommentGetCommentsByArticleId)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCommentsByArticleId(ctx, req.(*GetCommentsByArticleIdReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetCommentsByArticleIdResp)
		return ctx.Result(200, reply)
	}
}

func _Comment_GetLastSevenComment0_HTTP_Handler(srv CommentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetLastSevenCommentReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCommentGetLastSevenComment)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetLastSevenComment(ctx, req.(*GetLastSevenCommentReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetLastSevenCommentResp)
		return ctx.Result(200, reply)
	}
}

func _Comment_GetRandomComment0_HTTP_Handler(srv CommentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetRandomCommentReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCommentGetRandomComment)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetRandomComment(ctx, req.(*GetRandomCommentReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetRandomCommentResp)
		return ctx.Result(200, reply)
	}
}

type CommentHTTPClient interface {
	GetCommentsByArticleId(ctx context.Context, req *GetCommentsByArticleIdReq, opts ...http.CallOption) (rsp *GetCommentsByArticleIdResp, err error)
	GetLastSevenComment(ctx context.Context, req *GetLastSevenCommentReq, opts ...http.CallOption) (rsp *GetLastSevenCommentResp, err error)
	GetRandomComment(ctx context.Context, req *GetRandomCommentReq, opts ...http.CallOption) (rsp *GetRandomCommentResp, err error)
	PostComment(ctx context.Context, req *PostCommentReq, opts ...http.CallOption) (rsp *PostCommentResp, err error)
}

type CommentHTTPClientImpl struct {
	cc *http.Client
}

func NewCommentHTTPClient(client *http.Client) CommentHTTPClient {
	return &CommentHTTPClientImpl{client}
}

func (c *CommentHTTPClientImpl) GetCommentsByArticleId(ctx context.Context, in *GetCommentsByArticleIdReq, opts ...http.CallOption) (*GetCommentsByArticleIdResp, error) {
	var out GetCommentsByArticleIdResp
	pattern := "/api/comment/{article_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCommentGetCommentsByArticleId))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CommentHTTPClientImpl) GetLastSevenComment(ctx context.Context, in *GetLastSevenCommentReq, opts ...http.CallOption) (*GetLastSevenCommentResp, error) {
	var out GetLastSevenCommentResp
	pattern := "/api/comment/last_seven/{article_id}/{num}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCommentGetLastSevenComment))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CommentHTTPClientImpl) GetRandomComment(ctx context.Context, in *GetRandomCommentReq, opts ...http.CallOption) (*GetRandomCommentResp, error) {
	var out GetRandomCommentResp
	pattern := "/api/comment/random/{article_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCommentGetRandomComment))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CommentHTTPClientImpl) PostComment(ctx context.Context, in *PostCommentReq, opts ...http.CallOption) (*PostCommentResp, error) {
	var out PostCommentResp
	pattern := "/api/comment/post"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCommentPostComment))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

const OperationCodeProcessingGetCommitByJudgeId = "/service.problem.api.CodeProcessing/GetCommitByJudgeId"
const OperationCodeProcessingSubmitCode = "/service.problem.api.CodeProcessing/SubmitCode"

type CodeProcessingHTTPServer interface {
	GetCommitByJudgeId(context.Context, *GetCommitByJudgeIdReq) (*GetCommitByJudgeIdResp, error)
	SubmitCode(context.Context, *SubmitCodeReq) (*SubmitCodeResp, error)
}

func RegisterCodeProcessingHTTPServer(s *http.Server, srv CodeProcessingHTTPServer) {
	r := s.Route("/")
	r.POST("/api/code/post", _CodeProcessing_SubmitCode0_HTTP_Handler(srv))
	r.GET("/api/code/view-submission/{judge_id}", _CodeProcessing_GetCommitByJudgeId0_HTTP_Handler(srv))
}

func _CodeProcessing_SubmitCode0_HTTP_Handler(srv CodeProcessingHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SubmitCodeReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCodeProcessingSubmitCode)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SubmitCode(ctx, req.(*SubmitCodeReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SubmitCodeResp)
		return ctx.Result(200, reply)
	}
}

func _CodeProcessing_GetCommitByJudgeId0_HTTP_Handler(srv CodeProcessingHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCommitByJudgeIdReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCodeProcessingGetCommitByJudgeId)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCommitByJudgeId(ctx, req.(*GetCommitByJudgeIdReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetCommitByJudgeIdResp)
		return ctx.Result(200, reply)
	}
}

type CodeProcessingHTTPClient interface {
	GetCommitByJudgeId(ctx context.Context, req *GetCommitByJudgeIdReq, opts ...http.CallOption) (rsp *GetCommitByJudgeIdResp, err error)
	SubmitCode(ctx context.Context, req *SubmitCodeReq, opts ...http.CallOption) (rsp *SubmitCodeResp, err error)
}

type CodeProcessingHTTPClientImpl struct {
	cc *http.Client
}

func NewCodeProcessingHTTPClient(client *http.Client) CodeProcessingHTTPClient {
	return &CodeProcessingHTTPClientImpl{client}
}

func (c *CodeProcessingHTTPClientImpl) GetCommitByJudgeId(ctx context.Context, in *GetCommitByJudgeIdReq, opts ...http.CallOption) (*GetCommitByJudgeIdResp, error) {
	var out GetCommitByJudgeIdResp
	pattern := "/api/code/view-submission/{judge_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCodeProcessingGetCommitByJudgeId))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CodeProcessingHTTPClientImpl) SubmitCode(ctx context.Context, in *SubmitCodeReq, opts ...http.CallOption) (*SubmitCodeResp, error) {
	var out SubmitCodeResp
	pattern := "/api/code/post"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCodeProcessingSubmitCode))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
