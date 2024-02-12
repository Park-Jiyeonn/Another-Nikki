package service

import (
	"context"

	pb "Another-Nikki/interact_hub/service/api"
)

type ArticleService struct {
	pb.UnimplementedArticleServer
}

func NewArticleService() *ArticleService {
	return &ArticleService{}
}

func (s *ArticleService) PostArticle(ctx context.Context, req *pb.PostArticleReq) (*pb.PostArticleResp, error) {
	return &pb.PostArticleResp{}, nil
}
func (s *ArticleService) GetArticleById(ctx context.Context, req *pb.GetArticleByIdReq) (*pb.GetArticleByIdResp, error) {
	return &pb.GetArticleByIdResp{}, nil
}
func (s *ArticleService) GetArticleByPage(ctx context.Context, req *pb.GetArticleByPageReq) (*pb.GetArticleByPageResp, error) {
	return &pb.GetArticleByPageResp{}, nil
}
