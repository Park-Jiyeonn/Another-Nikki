package biz

import (
	"golang.org/x/net/context"
	"time"
)

type CommentRepo interface {
	PostComment(ctx context.Context, req *PostCommentReq) error
	GetCommentsByArticleId(ctx context.Context, req *GetCommentsByArticleIdReq) (*GetCommentsByArticleIdResp, error)
	GetLastSevenComment(ctx context.Context, req *GetLastSevenCommentReq) (*GetLastSevenCommentResp, error)
	GetCommentById(ctx context.Context, req *GetRandomCommentReq) (*GetRandomCommentResp, error)
	GetCommentSum(ctx context.Context) (sum int64, err error)
}

type PostCommentReq struct {
	Content      string `json:"content"`
	ArticleId    int64  `json:"article_id"`
	AuthorName   string `json:"author_name"`
	AuthorAvatar string `json:"author_avatar"`
	ParentId     int64  `json:"parent_id"`
	RootId       int64  `json:"root_id"`
}

type GetCommentsByArticleIdReq struct {
	ArticleId int64 `json:"article_id"`
}

type Comments struct {
	Content      string    `json:"content"`
	AuthorName   string    `json:"author_name"`
	AuthorAvatar string    `json:"author_avatar"`
	ParentId     int64     `json:"parent_id"`
	RootId       int64     `json:"root_id"`
	CreatedTime  time.Time `json:"created_time"`
}

type GetCommentsByArticleIdResp struct {
	Comments []*Comments `json:"comments"`
}

type GetLastSevenCommentReq struct {
	ArticleId int64 `json:"article_id"`
}

type GetLastSevenCommentResp struct {
	Comments []*Comments `json:"comments"`
}

type GetRandomCommentReq struct {
	ArticleId int64 `json:"article_id"`
	CommentId int64 `json:"comment_id"`
}

type GetRandomCommentResp struct {
	Comments *Comments `json:"comments"`
}
