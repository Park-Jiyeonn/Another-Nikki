package biz

import (
	"golang.org/x/net/context"
	"time"
)

type CommentRepo interface {
	PostComment(ctx context.Context, req *PostCommentReq) error
	GetCommentsByArticleId(ctx context.Context, req *GetCommentsByArticleIdReq) (*GetCommentsByArticleIdResp, error)
	GetLastSevenComment(ctx context.Context, req *GetLastSevenCommentReq) (*GetLastSevenCommentResp, error)
	GetCommentByOffset(ctx context.Context, req *GetRandomCommentReq) (*GetRandomCommentResp, error)
	GetCommentSum(ctx context.Context, req *GetCommentSumReq) (sum int64, err error)
}

type PostCommentReq struct {
	Content    string `db:"content"`
	ArticleId  int64  `db:"article_id"`
	Username   string `db:"username"`
	UserAvatar string `db:"user_avatar"`
	ParentId   int64  `db:"parent_id"`
	RootId     int64  `db:"root_id"`
	UserId     int64  `db:"user_id"`
	ParentName string `db:"parent_name"`
}

type GetCommentsByArticleIdReq struct {
	ArticleId int64 `db:"article_id"`
}

type Comments struct {
	CommentId   int64       `db:"comment_id"`
	Content     string      `db:"content"`
	Username    string      `db:"username"`
	UserAvatar  string      `db:"user_avatar"`
	ParentId    int64       `db:"parent_id"`
	RootId      int64       `db:"root_id"`
	CreatedTime time.Time   `db:"created_time"`
	Children    []*Comments `db:"-"`
	ParentName  string      `db:"parent_name"`
}

type GetCommentsByArticleIdResp struct {
	Comments []*Comments `db:"comments"`
}

type GetLastSevenCommentReq struct {
	ArticleId int64 `db:"article_id"`
	NumLimit  int64
}

type GetLastSevenCommentResp struct {
	Comments []*Comments `db:"comments"`
}

type GetRandomCommentReq struct {
	ArticleId     int64 `db:"article_id"`
	CommentOffset int64
}

type GetRandomCommentResp struct {
	Comments *Comments `db:"comments"`
}

type GetCommentSumReq struct {
	ArticleId int64 `db:"article_id"`
}
