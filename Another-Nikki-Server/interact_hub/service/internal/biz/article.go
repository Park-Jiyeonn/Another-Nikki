package biz

import (
	"golang.org/x/net/context"
	"time"
)

type ArticleRepo interface {
	PostArticle(ctx context.Context, req *PostArticleReq) (err error)
	GetArticleById(ctx context.Context, req *GetArticleByIdReq) (*GetArticleByIdResp, error)
	GetArticleByPage(ctx context.Context, req *GetArticleByPageReq) (*GetArticleByPageResp, error)
}

type PostArticleReq struct {
	ArticleTitle       string `json:"article_title"`
	ArticleDescription string `json:"article_description"`
	ArticleContent     string `json:"article_content"`
}

type GetArticleByIdReq struct {
	ArticleId int64 `json:"article_id"`
}

type GetArticleByIdResp struct {
	ArticleTitle       string    `json:"article_title"`
	ArticleDescription string    `json:"article_description"`
	ArticleContent     string    `json:"article_content"`
	CreateTime         time.Time `json:"create_time"`
}

type GetArticleByPageReq struct {
	PageNum  int64 `json:"page_num"`
	PageSize int64 `json:"page_size"`
}

type ArticlePageDetail struct {
	ArticleId    int64     `json:"article_id"`
	ArticleTitle string    `json:"article_title"`
	CreateTime   time.Time `json:"create_time"`
}

type GetArticleByPageResp struct {
	Articles []*ArticlePageDetail `json:"articles"`
}
