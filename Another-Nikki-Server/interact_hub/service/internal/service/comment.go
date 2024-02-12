package service

import (
	"Another-Nikki/interact_hub/service/internal/biz"
	"context"

	pb "Another-Nikki/interact_hub/service/api"
)

type CommentService struct {
	pb.UnimplementedCommentServer
	dao biz.CommentRepo
}

func NewCommentService(dao biz.CommentRepo) *CommentService {
	return &CommentService{
		dao: dao,
	}
}

func (s *CommentService) PostComment(ctx context.Context, req *pb.PostCommentReq) (*pb.PostCommentResp, error) {
	return &pb.PostCommentResp{}, nil
}
func (s *CommentService) GetCommentsByArticleId(ctx context.Context, req *pb.GetCommentsByArticleIdReq) (*pb.GetCommentsByArticleIdResp, error) {
	return &pb.GetCommentsByArticleIdResp{}, nil
}
func (s *CommentService) GetLastSevenComment(ctx context.Context, req *pb.GetLastSevenCommentReq) (*pb.GetLastSevenCommentResp, error) {
	return &pb.GetLastSevenCommentResp{}, nil
}
func (s *CommentService) GetRandomComment(ctx context.Context, req *pb.GetRandomCommentReq) (*pb.GetRandomCommentResp, error) {
	return &pb.GetRandomCommentResp{}, nil
}
