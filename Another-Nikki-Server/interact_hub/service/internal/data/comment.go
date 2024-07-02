package data

import (
	"Another-Nikki/interact_hub/service/internal/biz"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
	"sort"
)

type CommentServiceImpl struct {
	db *sqlx.DB
}

func NewCommentImpl(db *sqlx.DB) biz.CommentRepo {
	return &CommentServiceImpl{db: db}
}

// CREATE TABLE comments (
//    comment_id SERIAL PRIMARY KEY,
//    content VARCHAR(255) NOT NULL DEFAULT '',
//    article_id INT NOT NULL,
//	  user_id INT NOT NULL,
//    username VARCHAR(255) NOT NULL DEFAULT '',
//    user_avatar VARCHAR(255) DEFAULT '',
//    parent_id INT,
//    parent_name VARCHAR(255) DEFAULT '',
//    root_id INT,
//    created_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//    updated_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
// );

func (s *CommentServiceImpl) PostComment(ctx context.Context, req *biz.PostCommentReq) (err error) {
	_, err = s.db.ExecContext(ctx, "INSERT INTO comments (content, article_id, username, user_avatar, parent_id, root_id, user_id, parent_name) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		req.Content, req.ArticleId, req.Username, req.UserAvatar, req.ParentId, req.RootId, req.UserId, req.ParentName)
	return
}

func (s *CommentServiceImpl) getChildren(ctx context.Context, articleId, rootId int64) (resp []*biz.Comments, err error) {
	resp = make([]*biz.Comments, 0)
	rows, err := s.db.QueryxContext(ctx, "SELECT content, username, user_avatar, parent_id, root_id, created_time, parent_name FROM comments WHERE article_id = ? and root_id = ?", articleId, rootId)
	for rows.Next() {
		var comment biz.Comments
		err = rows.StructScan(&comment)
		if err != nil {
			return nil, err
		}
		resp = append(resp, &comment)
	}
	return
}
func (s *CommentServiceImpl) GetCommentsByArticleId(ctx context.Context, req *biz.GetCommentsByArticleIdReq) (*biz.GetCommentsByArticleIdResp, error) {
	rows, err := s.db.QueryxContext(ctx, "SELECT comment_id, content, username, user_avatar, parent_id, root_id, created_time FROM comments WHERE article_id = ? and root_id = 0", req.ArticleId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*biz.Comments
	for rows.Next() {
		var comment biz.Comments
		err = rows.StructScan(&comment)
		if err != nil {
			return nil, err
		}
		children, err := s.getChildren(ctx, req.ArticleId, comment.CommentId)
		if err != nil {
			return nil, err
		}
		comment.Children = append(comment.Children, children...)
		comments = append(comments, &comment)
	}

	sort.Slice(comments, func(i, j int) bool {
		return comments[i].CreatedTime.Unix() < comments[j].CreatedTime.Unix()
	})
	return &biz.GetCommentsByArticleIdResp{Comments: comments}, nil
}

func (s *CommentServiceImpl) GetLastSevenComment(ctx context.Context, req *biz.GetLastSevenCommentReq) (*biz.GetLastSevenCommentResp, error) {
	rows, err := s.db.QueryxContext(ctx, "SELECT comment_id, content, username, user_avatar, parent_id, root_id, created_time FROM comments WHERE article_id = ? and root_id = 0 ORDER BY created_time DESC LIMIT ?", req.ArticleId, req.NumLimit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*biz.Comments
	for rows.Next() {
		var comment biz.Comments
		err = rows.StructScan(&comment)
		if err != nil {
			return nil, err
		}
		children, err := s.getChildren(ctx, req.ArticleId, comment.CommentId)
		if err != nil {
			return nil, err
		}
		comment.Children = append(comment.Children, children...)
		comments = append(comments, &comment)
	}

	sort.Slice(comments, func(i, j int) bool {
		return comments[i].CreatedTime.Unix() < comments[j].CreatedTime.Unix()
	})
	return &biz.GetLastSevenCommentResp{Comments: comments}, nil
}

func (s *CommentServiceImpl) GetCommentByOffset(ctx context.Context, req *biz.GetRandomCommentReq) (*biz.GetRandomCommentResp, error) {
	var comment biz.Comments
	err := s.db.GetContext(ctx, &comment, "SELECT content, username, user_avatar, parent_id, root_id, created_time FROM comments WHERE article_id = ? ORDER BY comment_id LIMIT 1 OFFSET ?",
		req.ArticleId,
		req.CommentOffset,
	)
	if err != nil {
		return nil, err
	}

	return &biz.GetRandomCommentResp{Comments: &comment}, nil
}

func (s *CommentServiceImpl) GetCommentSum(ctx context.Context, req *biz.GetCommentSumReq) (sum int64, err error) {
	err = s.db.QueryRowContext(ctx, "SELECT COUNT(comment_id) from comments where article_id = ?", req.ArticleId).Scan(&sum)
	return
}
