package data

import (
	"Another-Nikki/interact_hub/service/internal/biz"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
)

type CommentServiceImpl struct {
	db *sqlx.DB
}

func NewCommentImpl(db *sqlx.DB) biz.CommentRepo {
	return &CommentServiceImpl{db: db}
}

// CREATE TABLE comments (
//    comment_id SERIAL PRIMARY KEY,
//    content VARCHAR(255) NOT NULL,
//    article_id INT NOT NULL,
//    author_name VARCHAR(255) NOT NULL,
//    author_avatar VARCHAR(255),
//    parent_id INT,
//    root_id INT,
//    created_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//    updated_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
// );

func (s *CommentServiceImpl) PostComment(ctx context.Context, req *biz.PostCommentReq) (err error) {
	_, err = s.db.ExecContext(ctx, "INSERT INTO comments (content, article_id, author_name, author_avatar, parent_id, root_id) VALUES (?, ?, ?, ?, ?, ?)",
		req.Content, req.ArticleId, req.AuthorName, req.AuthorAvatar, req.ParentId, req.RootId)
	return
}

func (s *CommentServiceImpl) GetCommentsByArticleId(ctx context.Context, req *biz.GetCommentsByArticleIdReq) (*biz.GetCommentsByArticleIdResp, error) {
	rows, err := s.db.QueryxContext(ctx, "SELECT content, author_name, author_avatar, parent_id, root_id, created_time FROM comments WHERE article_id = ?", req.ArticleId)
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
		comments = append(comments, &comment)
	}

	return &biz.GetCommentsByArticleIdResp{Comments: comments}, nil
}

func (s *CommentServiceImpl) GetLastSevenComment(ctx context.Context, req *biz.GetLastSevenCommentReq) (*biz.GetLastSevenCommentResp, error) {
	rows, err := s.db.QueryxContext(ctx, "SELECT content, author_name, author_avatar, parent_id, root_id FROM comments WHERE article_id = ? ORDER BY created_time DESC LIMIT 7", req.ArticleId)
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
		comments = append(comments, &comment)
	}

	return &biz.GetLastSevenCommentResp{Comments: comments}, nil
}

func (s *CommentServiceImpl) GetCommentById(ctx context.Context, req *biz.GetRandomCommentReq) (*biz.GetRandomCommentResp, error) {
	var comment biz.Comments
	err := s.db.GetContext(ctx, &comment, "SELECT content, author_name, author_avatar, parent_id, root_id FROM comments WHERE article_id = ? and comment_id = ?",
		req.ArticleId,
		req.CommentId,
	)
	if err != nil {
		return nil, err
	}

	return &biz.GetRandomCommentResp{Comments: &comment}, nil
}

func (s *CommentServiceImpl) GetCommentSum(ctx context.Context) (sum int64, err error) {
	err = s.db.QueryRowContext(ctx, "SELECT COUNT(comment_id) from comments").Scan(&sum)
	return
}
