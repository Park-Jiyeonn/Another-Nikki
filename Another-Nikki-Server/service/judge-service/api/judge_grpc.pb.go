// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: judge.proto

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
	Judge_Judge_FullMethodName     = "/AnotherNikki.oj.judge.judge/Judge"
	Judge_OnlineRun_FullMethodName = "/AnotherNikki.oj.judge.judge/OnlineRun"
)

// JudgeClient is the client API for Judge service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JudgeClient interface {
	Judge(ctx context.Context, in *JudgeReq, opts ...grpc.CallOption) (*JudgeResp, error)
	OnlineRun(ctx context.Context, in *OnlineRunReq, opts ...grpc.CallOption) (*OnlineRunResp, error)
}

type judgeClient struct {
	cc grpc.ClientConnInterface
}

func NewJudgeClient(cc grpc.ClientConnInterface) JudgeClient {
	return &judgeClient{cc}
}

func (c *judgeClient) Judge(ctx context.Context, in *JudgeReq, opts ...grpc.CallOption) (*JudgeResp, error) {
	out := new(JudgeResp)
	err := c.cc.Invoke(ctx, Judge_Judge_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *judgeClient) OnlineRun(ctx context.Context, in *OnlineRunReq, opts ...grpc.CallOption) (*OnlineRunResp, error) {
	out := new(OnlineRunResp)
	err := c.cc.Invoke(ctx, Judge_OnlineRun_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JudgeServer is the server API for Judge service.
// All implementations must embed UnimplementedJudgeServer
// for forward compatibility
type JudgeServer interface {
	Judge(context.Context, *JudgeReq) (*JudgeResp, error)
	OnlineRun(context.Context, *OnlineRunReq) (*OnlineRunResp, error)
	mustEmbedUnimplementedJudgeServer()
}

// UnimplementedJudgeServer must be embedded to have forward compatible implementations.
type UnimplementedJudgeServer struct {
}

func (UnimplementedJudgeServer) Judge(context.Context, *JudgeReq) (*JudgeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Judge not implemented")
}
func (UnimplementedJudgeServer) OnlineRun(context.Context, *OnlineRunReq) (*OnlineRunResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnlineRun not implemented")
}
func (UnimplementedJudgeServer) mustEmbedUnimplementedJudgeServer() {}

// UnsafeJudgeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JudgeServer will
// result in compilation errors.
type UnsafeJudgeServer interface {
	mustEmbedUnimplementedJudgeServer()
}

func RegisterJudgeServer(s grpc.ServiceRegistrar, srv JudgeServer) {
	s.RegisterService(&Judge_ServiceDesc, srv)
}

func _Judge_Judge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JudgeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JudgeServer).Judge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Judge_Judge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JudgeServer).Judge(ctx, req.(*JudgeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Judge_OnlineRun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnlineRunReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JudgeServer).OnlineRun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Judge_OnlineRun_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JudgeServer).OnlineRun(ctx, req.(*OnlineRunReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Judge_ServiceDesc is the grpc.ServiceDesc for Judge service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Judge_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AnotherNikki.oj.judge.judge",
	HandlerType: (*JudgeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Judge",
			Handler:    _Judge_Judge_Handler,
		},
		{
			MethodName: "OnlineRun",
			Handler:    _Judge_OnlineRun_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "judge.proto",
}
