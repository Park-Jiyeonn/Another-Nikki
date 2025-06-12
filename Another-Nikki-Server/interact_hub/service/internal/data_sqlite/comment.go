package data_sqlite

import (
	"Another-Nikki/interact_hub/service/internal/biz"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
	"sort"
)

type CommentServiceImpl struct {
	db *sqlx.DB
}

func InitCommentTable(db *sqlx.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS comments (
		comment_id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT NOT NULL,
		article_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		username TEXT NOT NULL DEFAULT '',
		user_avatar TEXT DEFAULT '',
		parent_id INTEGER,
		parent_name TEXT DEFAULT '',
		root_id INTEGER,
		created_time DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_time DATETIME DEFAULT CURRENT_TIMESTAMP
	);`
	_, err := db.Exec(schema)
	return err
}

func NewCommentImpl(data *Data) biz.CommentRepo {
	if err := InitCommentTable(data.GlobalDB); err != nil {
		panic(err)
	}
	return &CommentServiceImpl{db: data.GlobalDB}
}

func (s *CommentServiceImpl) PostComment(ctx context.Context, req *biz.PostCommentReq) (err error) {
	_, err = s.db.ExecContext(ctx,
		`INSERT INTO comments 
		(content, article_id, username, user_avatar, parent_id, root_id, user_id, parent_name, updated_time)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP)`,
		req.Content, req.ArticleId, req.Username, req.UserAvatar,
		req.ParentId, req.RootId, req.UserId, req.ParentName)
	return
}

func (s *CommentServiceImpl) getChildren(ctx context.Context, articleId, rootId int64) (resp []*biz.Comments, err error) {
	resp = make([]*biz.Comments, 0)
	rows, err := s.db.QueryxContext(ctx,
		`SELECT comment_id, content, username, user_avatar, parent_id, root_id, created_time, parent_name 
		FROM comments 
		WHERE article_id = ? AND root_id = ?`, articleId, rootId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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
	rows, err := s.db.QueryxContext(ctx,
		`SELECT comment_id, content, username, user_avatar, parent_id, root_id, created_time 
		FROM comments 
		WHERE article_id = ? AND root_id = 0`, req.ArticleId)
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
	rows, err := s.db.QueryxContext(ctx,
		`SELECT comment_id, content, username, user_avatar, parent_id, root_id, created_time 
		FROM comments 
		WHERE article_id = ? AND root_id = 0 
		ORDER BY created_time DESC 
		LIMIT ?`, req.ArticleId, req.NumLimit)
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
	err := s.db.GetContext(ctx, &comment,
		`SELECT comment_id, content, username, user_avatar, parent_id, root_id, created_time 
		FROM comments 
		WHERE article_id = ? 
		ORDER BY comment_id 
		LIMIT 1 OFFSET ?`, req.ArticleId, req.CommentOffset)
	if err != nil {
		return nil, err
	}

	return &biz.GetRandomCommentResp{Comments: &comment}, nil
}

func (s *CommentServiceImpl) GetCommentSum(ctx context.Context, req *biz.GetCommentSumReq) (sum int64, err error) {
	err = s.db.QueryRowContext(ctx,
		`SELECT COUNT(comment_id) FROM comments WHERE article_id = ?`, req.ArticleId).Scan(&sum)
	return
}
