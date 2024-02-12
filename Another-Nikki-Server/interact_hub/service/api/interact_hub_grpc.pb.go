// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: interact_hub.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Problem_PostProblem_FullMethodName      = "/service.problem.api.Problem/PostProblem"
	Problem_GetProblemById_FullMethodName   = "/service.problem.api.Problem/GetProblemById"
	Problem_GetProblemByPage_FullMethodName = "/service.problem.api.Problem/GetProblemByPage"
)

// ProblemClient is the client API for Problem service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProblemClient interface {
	PostProblem(ctx context.Context, in *PostProblemReq, opts ...grpc.CallOption) (*PostProblemResp, error)
	GetProblemById(ctx context.Context, in *GetProblemByIdReq, opts ...grpc.CallOption) (*GetProblemByIdResp, error)
	GetProblemByPage(ctx context.Context, in *GetProblemByPageReq, opts ...grpc.CallOption) (*GetProblemByPageResp, error)
}

type problemClient struct {
	cc grpc.ClientConnInterface
}

func NewProblemClient(cc grpc.ClientConnInterface) ProblemClient {
	return &problemClient{cc}
}

func (c *problemClient) PostProblem(ctx context.Context, in *PostProblemReq, opts ...grpc.CallOption) (*PostProblemResp, error) {
	out := new(PostProblemResp)
	err := c.cc.Invoke(ctx, Problem_PostProblem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemClient) GetProblemById(ctx context.Context, in *GetProblemByIdReq, opts ...grpc.CallOption) (*GetProblemByIdResp, error) {
	out := new(GetProblemByIdResp)
	err := c.cc.Invoke(ctx, Problem_GetProblemById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemClient) GetProblemByPage(ctx context.Context, in *GetProblemByPageReq, opts ...grpc.CallOption) (*GetProblemByPageResp, error) {
	out := new(GetProblemByPageResp)
	err := c.cc.Invoke(ctx, Problem_GetProblemByPage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProblemServer is the server API for Problem service.
// All implementations must embed UnimplementedProblemServer
// for forward compatibility
type ProblemServer interface {
	PostProblem(context.Context, *PostProblemReq) (*PostProblemResp, error)
	GetProblemById(context.Context, *GetProblemByIdReq) (*GetProblemByIdResp, error)
	GetProblemByPage(context.Context, *GetProblemByPageReq) (*GetProblemByPageResp, error)
	mustEmbedUnimplementedProblemServer()
}

// UnimplementedProblemServer must be embedded to have forward compatible implementations.
type UnimplementedProblemServer struct {
}

func (UnimplementedProblemServer) PostProblem(context.Context, *PostProblemReq) (*PostProblemResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostProblem not implemented")
}
func (UnimplementedProblemServer) GetProblemById(context.Context, *GetProblemByIdReq) (*GetProblemByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProblemById not implemented")
}
func (UnimplementedProblemServer) GetProblemByPage(context.Context, *GetProblemByPageReq) (*GetProblemByPageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProblemByPage not implemented")
}
func (UnimplementedProblemServer) mustEmbedUnimplementedProblemServer() {}

// UnsafeProblemServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProblemServer will
// result in compilation errors.
type UnsafeProblemServer interface {
	mustEmbedUnimplementedProblemServer()
}

func RegisterProblemServer(s grpc.ServiceRegistrar, srv ProblemServer) {
	s.RegisterService(&Problem_ServiceDesc, srv)
}

func _Problem_PostProblem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostProblemReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServer).PostProblem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Problem_PostProblem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServer).PostProblem(ctx, req.(*PostProblemReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Problem_GetProblemById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProblemByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServer).GetProblemById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Problem_GetProblemById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServer).GetProblemById(ctx, req.(*GetProblemByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Problem_GetProblemByPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProblemByPageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServer).GetProblemByPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Problem_GetProblemByPage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServer).GetProblemByPage(ctx, req.(*GetProblemByPageReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Problem_ServiceDesc is the grpc.ServiceDesc for Problem service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Problem_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.problem.api.Problem",
	HandlerType: (*ProblemServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostProblem",
			Handler:    _Problem_PostProblem_Handler,
		},
		{
			MethodName: "GetProblemById",
			Handler:    _Problem_GetProblemById_Handler,
		},
		{
			MethodName: "GetProblemByPage",
			Handler:    _Problem_GetProblemByPage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interact_hub.proto",
}

const (
	Article_PostArticle_FullMethodName      = "/service.problem.api.Article/PostArticle"
	Article_GetArticleById_FullMethodName   = "/service.problem.api.Article/GetArticleById"
	Article_GetArticleByPage_FullMethodName = "/service.problem.api.Article/GetArticleByPage"
)

// ArticleClient is the client API for Article service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArticleClient interface {
	PostArticle(ctx context.Context, in *PostArticleReq, opts ...grpc.CallOption) (*PostArticleResp, error)
	GetArticleById(ctx context.Context, in *GetArticleByIdReq, opts ...grpc.CallOption) (*GetArticleByIdResp, error)
	GetArticleByPage(ctx context.Context, in *GetArticleByPageReq, opts ...grpc.CallOption) (*GetArticleByPageResp, error)
}

type articleClient struct {
	cc grpc.ClientConnInterface
}

func NewArticleClient(cc grpc.ClientConnInterface) ArticleClient {
	return &articleClient{cc}
}

func (c *articleClient) PostArticle(ctx context.Context, in *PostArticleReq, opts ...grpc.CallOption) (*PostArticleResp, error) {
	out := new(PostArticleResp)
	err := c.cc.Invoke(ctx, Article_PostArticle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleClient) GetArticleById(ctx context.Context, in *GetArticleByIdReq, opts ...grpc.CallOption) (*GetArticleByIdResp, error) {
	out := new(GetArticleByIdResp)
	err := c.cc.Invoke(ctx, Article_GetArticleById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleClient) GetArticleByPage(ctx context.Context, in *GetArticleByPageReq, opts ...grpc.CallOption) (*GetArticleByPageResp, error) {
	out := new(GetArticleByPageResp)
	err := c.cc.Invoke(ctx, Article_GetArticleByPage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticleServer is the server API for Article service.
// All implementations must embed UnimplementedArticleServer
// for forward compatibility
type ArticleServer interface {
	PostArticle(context.Context, *PostArticleReq) (*PostArticleResp, error)
	GetArticleById(context.Context, *GetArticleByIdReq) (*GetArticleByIdResp, error)
	GetArticleByPage(context.Context, *GetArticleByPageReq) (*GetArticleByPageResp, error)
	mustEmbedUnimplementedArticleServer()
}

// UnimplementedArticleServer must be embedded to have forward compatible implementations.
type UnimplementedArticleServer struct {
}

func (UnimplementedArticleServer) PostArticle(context.Context, *PostArticleReq) (*PostArticleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostArticle not implemented")
}
func (UnimplementedArticleServer) GetArticleById(context.Context, *GetArticleByIdReq) (*GetArticleByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticleById not implemented")
}
func (UnimplementedArticleServer) GetArticleByPage(context.Context, *GetArticleByPageReq) (*GetArticleByPageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticleByPage not implemented")
}
func (UnimplementedArticleServer) mustEmbedUnimplementedArticleServer() {}

// UnsafeArticleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArticleServer will
// result in compilation errors.
type UnsafeArticleServer interface {
	mustEmbedUnimplementedArticleServer()
}

func RegisterArticleServer(s grpc.ServiceRegistrar, srv ArticleServer) {
	s.RegisterService(&Article_ServiceDesc, srv)
}

func _Article_PostArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostArticleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).PostArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Article_PostArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).PostArticle(ctx, req.(*PostArticleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Article_GetArticleById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).GetArticleById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Article_GetArticleById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).GetArticleById(ctx, req.(*GetArticleByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Article_GetArticleByPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleByPageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).GetArticleByPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Article_GetArticleByPage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).GetArticleByPage(ctx, req.(*GetArticleByPageReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Article_ServiceDesc is the grpc.ServiceDesc for Article service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Article_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.problem.api.Article",
	HandlerType: (*ArticleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostArticle",
			Handler:    _Article_PostArticle_Handler,
		},
		{
			MethodName: "GetArticleById",
			Handler:    _Article_GetArticleById_Handler,
		},
		{
			MethodName: "GetArticleByPage",
			Handler:    _Article_GetArticleByPage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interact_hub.proto",
}

const (
	User_Login_FullMethodName             = "/service.problem.api.User/Login"
	User_Register_FullMethodName          = "/service.problem.api.User/Register"
	User_GetUserByUserName_FullMethodName = "/service.problem.api.User/GetUserByUserName"
	User_GetUserById_FullMethodName       = "/service.problem.api.User/GetUserById"
)

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
	GetUserByUserName(ctx context.Context, in *GetUserByUserNameReq, opts ...grpc.CallOption) (*GetUserByUserNameResp, error)
	GetUserById(ctx context.Context, in *GetUserByIdReq, opts ...grpc.CallOption) (*GetUserByIdResp, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, User_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	out := new(RegisterResp)
	err := c.cc.Invoke(ctx, User_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserByUserName(ctx context.Context, in *GetUserByUserNameReq, opts ...grpc.CallOption) (*GetUserByUserNameResp, error) {
	out := new(GetUserByUserNameResp)
	err := c.cc.Invoke(ctx, User_GetUserByUserName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserById(ctx context.Context, in *GetUserByIdReq, opts ...grpc.CallOption) (*GetUserByIdResp, error) {
	out := new(GetUserByIdResp)
	err := c.cc.Invoke(ctx, User_GetUserById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	Login(context.Context, *LoginReq) (*LoginResp, error)
	Register(context.Context, *RegisterReq) (*RegisterResp, error)
	GetUserByUserName(context.Context, *GetUserByUserNameReq) (*GetUserByUserNameResp, error)
	GetUserById(context.Context, *GetUserByIdReq) (*GetUserByIdResp, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) Login(context.Context, *LoginReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServer) Register(context.Context, *RegisterReq) (*RegisterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserServer) GetUserByUserName(context.Context, *GetUserByUserNameReq) (*GetUserByUserNameResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByUserName not implemented")
}
func (UnimplementedUserServer) GetUserById(context.Context, *GetUserByIdReq) (*GetUserByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserByUserName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByUserNameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserByUserName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetUserByUserName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserByUserName(ctx, req.(*GetUserByUserNameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetUserById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserById(ctx, req.(*GetUserByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.problem.api.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _User_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _User_Register_Handler,
		},
		{
			MethodName: "GetUserByUserName",
			Handler:    _User_GetUserByUserName_Handler,
		},
		{
			MethodName: "GetUserById",
			Handler:    _User_GetUserById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interact_hub.proto",
}

const (
	Comment_PostComment_FullMethodName            = "/service.problem.api.Comment/PostComment"
	Comment_GetCommentsByArticleId_FullMethodName = "/service.problem.api.Comment/GetCommentsByArticleId"
	Comment_GetLastSevenComment_FullMethodName    = "/service.problem.api.Comment/GetLastSevenComment"
	Comment_GetRandomComment_FullMethodName       = "/service.problem.api.Comment/GetRandomComment"
)

// CommentClient is the client API for Comment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommentClient interface {
	PostComment(ctx context.Context, in *PostCommentReq, opts ...grpc.CallOption) (*PostCommentResp, error)
	GetCommentsByArticleId(ctx context.Context, in *GetCommentsByArticleIdReq, opts ...grpc.CallOption) (*GetCommentsByArticleIdResp, error)
	GetLastSevenComment(ctx context.Context, in *GetLastSevenCommentReq, opts ...grpc.CallOption) (*GetLastSevenCommentResp, error)
	GetRandomComment(ctx context.Context, in *GetRandomCommentReq, opts ...grpc.CallOption) (*GetRandomCommentResp, error)
}

type commentClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentClient(cc grpc.ClientConnInterface) CommentClient {
	return &commentClient{cc}
}

func (c *commentClient) PostComment(ctx context.Context, in *PostCommentReq, opts ...grpc.CallOption) (*PostCommentResp, error) {
	out := new(PostCommentResp)
	err := c.cc.Invoke(ctx, Comment_PostComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) GetCommentsByArticleId(ctx context.Context, in *GetCommentsByArticleIdReq, opts ...grpc.CallOption) (*GetCommentsByArticleIdResp, error) {
	out := new(GetCommentsByArticleIdResp)
	err := c.cc.Invoke(ctx, Comment_GetCommentsByArticleId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) GetLastSevenComment(ctx context.Context, in *GetLastSevenCommentReq, opts ...grpc.CallOption) (*GetLastSevenCommentResp, error) {
	out := new(GetLastSevenCommentResp)
	err := c.cc.Invoke(ctx, Comment_GetLastSevenComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentClient) GetRandomComment(ctx context.Context, in *GetRandomCommentReq, opts ...grpc.CallOption) (*GetRandomCommentResp, error) {
	out := new(GetRandomCommentResp)
	err := c.cc.Invoke(ctx, Comment_GetRandomComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentServer is the server API for Comment service.
// All implementations must embed UnimplementedCommentServer
// for forward compatibility
type CommentServer interface {
	PostComment(context.Context, *PostCommentReq) (*PostCommentResp, error)
	GetCommentsByArticleId(context.Context, *GetCommentsByArticleIdReq) (*GetCommentsByArticleIdResp, error)
	GetLastSevenComment(context.Context, *GetLastSevenCommentReq) (*GetLastSevenCommentResp, error)
	GetRandomComment(context.Context, *GetRandomCommentReq) (*GetRandomCommentResp, error)
	mustEmbedUnimplementedCommentServer()
}

// UnimplementedCommentServer must be embedded to have forward compatible implementations.
type UnimplementedCommentServer struct {
}

func (UnimplementedCommentServer) PostComment(context.Context, *PostCommentReq) (*PostCommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostComment not implemented")
}
func (UnimplementedCommentServer) GetCommentsByArticleId(context.Context, *GetCommentsByArticleIdReq) (*GetCommentsByArticleIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentsByArticleId not implemented")
}
func (UnimplementedCommentServer) GetLastSevenComment(context.Context, *GetLastSevenCommentReq) (*GetLastSevenCommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLastSevenComment not implemented")
}
func (UnimplementedCommentServer) GetRandomComment(context.Context, *GetRandomCommentReq) (*GetRandomCommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRandomComment not implemented")
}
func (UnimplementedCommentServer) mustEmbedUnimplementedCommentServer() {}

// UnsafeCommentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentServer will
// result in compilation errors.
type UnsafeCommentServer interface {
	mustEmbedUnimplementedCommentServer()
}

func RegisterCommentServer(s grpc.ServiceRegistrar, srv CommentServer) {
	s.RegisterService(&Comment_ServiceDesc, srv)
}

func _Comment_PostComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).PostComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Comment_PostComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).PostComment(ctx, req.(*PostCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_GetCommentsByArticleId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentsByArticleIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).GetCommentsByArticleId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Comment_GetCommentsByArticleId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).GetCommentsByArticleId(ctx, req.(*GetCommentsByArticleIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_GetLastSevenComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLastSevenCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).GetLastSevenComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Comment_GetLastSevenComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).GetLastSevenComment(ctx, req.(*GetLastSevenCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comment_GetRandomComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRandomCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServer).GetRandomComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Comment_GetRandomComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServer).GetRandomComment(ctx, req.(*GetRandomCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Comment_ServiceDesc is the grpc.ServiceDesc for Comment service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Comment_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.problem.api.Comment",
	HandlerType: (*CommentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostComment",
			Handler:    _Comment_PostComment_Handler,
		},
		{
			MethodName: "GetCommentsByArticleId",
			Handler:    _Comment_GetCommentsByArticleId_Handler,
		},
		{
			MethodName: "GetLastSevenComment",
			Handler:    _Comment_GetLastSevenComment_Handler,
		},
		{
			MethodName: "GetRandomComment",
			Handler:    _Comment_GetRandomComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interact_hub.proto",
}