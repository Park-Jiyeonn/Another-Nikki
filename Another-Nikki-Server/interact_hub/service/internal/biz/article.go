package biz

import "golang.org/x/net/context"

type Articles struct {
	ArticleId    int64
	ArticleTitle string
	CreateTime   string
}
type ArticleRepo interface {
	CreateArticle(ctx context.Context, articleId int64, title, description, content string) (err error)
	GetArticleById(ctx context.Context, articleId int64) (articleTitle, articleDescription, articleContent, createTime string, err error)
	GetArticleByPage(ctx context.Context, pageNum, pageSize int64) (articles []*Articles, err error)
}
