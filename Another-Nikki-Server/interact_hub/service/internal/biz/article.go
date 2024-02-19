package biz

import (
	"golang.org/x/net/context"
	"time"
)

type ArticleRepo interface {
	PostArticle(ctx context.Context, req *PostArticleReq) (err error)
	UpdateArticle(ctx context.Context, req *UpdateArticleReq) (err error)
	GetArticleById(ctx context.Context, req *GetArticleByIdReq) (*GetArticleByIdResp, error)
	GetArticleByPage(ctx context.Context, req *GetArticleByPageReq) (*GetArticleByPageResp, error)
}

type PostArticleReq struct {
	ArticleTitle       string `db:"article_title"`
	ArticleDescription string `db:"article_description"`
	ArticleContent     string `db:"article_content"`
}

type UpdateArticleReq struct {
	ArticleId          int64  `db:"article_id"`
	ArticleTitle       string `db:"article_title"`
	ArticleDescription string `db:"article_description"`
	ArticleContent     string `db:"article_content"`
}

type GetArticleByIdReq struct {
	ArticleId int64 `db:"article_id"`
}

type GetArticleByIdResp struct {
	ArticleTitle       string    `db:"article_title"`
	ArticleDescription string    `db:"article_description"`
	ArticleContent     string    `db:"article_content"`
	CreatedTime        time.Time `db:"created_time"`
}

type GetArticleByPageReq struct {
	PageNum  int64 `db:"page_num"`
	PageSize int64 `db:"page_size"`
}

type ArticlePageDetail struct {
	ArticleId          int64     `db:"article_id"`
	ArticleTitle       string    `db:"article_title"`
	CreatedTime        time.Time `db:"created_time"`
	ArticleDescription string    `db:"article_description"`
}

type GetArticleByPageResp struct {
	Articles []*ArticlePageDetail `db:"articles"`
}
