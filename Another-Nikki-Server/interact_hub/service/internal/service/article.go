package service

import (
	"Another-Nikki/interact_hub/service/internal/biz"
	"context"
	"time"

	pb "Another-Nikki/interact_hub/service/api"
)

type ArticleService struct {
	pb.UnimplementedArticleServer
	dao biz.ArticleRepo
}

func NewArticleService(dao biz.ArticleRepo) *ArticleService {
	return &ArticleService{
		dao: dao,
	}
}

func (s *ArticleService) PostArticle(ctx context.Context, req *pb.PostArticleReq) (resp *pb.PostArticleResp, err error) {
	resp = new(pb.PostArticleResp)
	if req.ArticleId != 0 {
		err = s.dao.UpdateArticle(ctx, &biz.UpdateArticleReq{
			ArticleId:          req.ArticleId,
			ArticleTitle:       req.ArticleTitle,
			ArticleDescription: req.ArticleDescription,
			ArticleContent:     req.ArticleContent,
		})
		return
	}
	err = s.dao.PostArticle(ctx, &biz.PostArticleReq{
		ArticleTitle:       req.ArticleTitle,
		ArticleDescription: req.ArticleDescription,
		ArticleContent:     req.ArticleContent,
	})
	return
}
func (s *ArticleService) GetArticleById(ctx context.Context, req *pb.GetArticleByIdReq) (*pb.GetArticleByIdResp, error) {
	resp := new(pb.GetArticleByIdResp)
	article, err := s.dao.GetArticleById(ctx, &biz.GetArticleByIdReq{
		ArticleId: req.ArticleId,
	})
	if err != nil {
		return nil, err
	}
	resp.ArticleTitle = article.ArticleTitle
	resp.ArticleDescription = article.ArticleDescription
	resp.CreateTime = article.CreatedTime.Format(time.DateTime)
	resp.ArticleContent = article.ArticleContent
	return resp, nil
}
func (s *ArticleService) GetArticleByPage(ctx context.Context, req *pb.GetArticleByPageReq) (resp *pb.GetArticleByPageResp, err error) {
	resp = new(pb.GetArticleByPageResp)
	articles, err := s.dao.GetArticleByPage(ctx, &biz.GetArticleByPageReq{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	resp.Articles = make([]*pb.ArticlePageDetail, 0, len(articles.Articles))
	for _, val := range articles.Articles {
		resp.Articles = append(resp.Articles, &pb.ArticlePageDetail{
			ArticleId:    val.ArticleId,
			ArticleTitle: val.ArticleTitle,
			CreateTime:   val.CreatedTime.Format(time.DateTime),
		})
	}
	return
}
