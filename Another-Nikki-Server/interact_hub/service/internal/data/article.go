package data

import (
	"Another-Nikki/interact_hub/service/internal/biz"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
	"time"
)

type ArticleDataMysql struct {
	ProblemId int64 `db:"problem_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type articleImpl struct {
	db *sqlx.DB
}

func NewArticleRepo(data *Data) biz.ArticleRepo {
	return &articleImpl{
		db: data.GlobalDB,
	}
}

// CREATE TABLE articles (
//    article_id SERIAL PRIMARY KEY,
//    article_title VARCHAR(255) NOT NULL,
//    article_description VARCHAR(255),
//    article_content TEXT,
//    created_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//    updated_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
// );

func (s *articleImpl) PostArticle(ctx context.Context, req *biz.PostArticleReq) (err error) {
	_, err = s.db.ExecContext(ctx, "INSERT INTO articles (article_title, article_description, article_content) VALUES (?, ?, ?)",
		req.ArticleTitle, req.ArticleDescription, req.ArticleContent)
	return
}
func (s *articleImpl) UpdateArticle(ctx context.Context, req *biz.UpdateArticleReq) (err error) {
	sqlStr := "UPDATE articles set article_title = ?, article_description = ?, article_content = ? where article_id = ?"
	_, err = s.db.ExecContext(ctx, sqlStr, req.ArticleTitle, req.ArticleDescription, req.ArticleContent, req.ArticleId)
	return
}
func (s *articleImpl) GetArticleById(ctx context.Context, req *biz.GetArticleByIdReq) (resp *biz.GetArticleByIdResp, err error) {
	resp = new(biz.GetArticleByIdResp)
	err = s.db.GetContext(ctx, resp, "SELECT article_title, article_description, article_content, created_time FROM articles WHERE article_id = ?", req.ArticleId)
	return
}
func (s *articleImpl) GetArticleByPage(ctx context.Context, req *biz.GetArticleByPageReq) (resp *biz.GetArticleByPageResp, err error) {
	rows, err := s.db.QueryxContext(ctx, "SELECT article_id, article_title, created_time FROM articles ORDER BY created_time DESC LIMIT ? OFFSET ?",
		req.PageSize, (req.PageNum-1)*req.PageSize)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var article biz.ArticlePageDetail
		err = rows.StructScan(&article)
		if err != nil {
			return nil, err
		}
		resp.Articles = append(resp.Articles, &article)
	}
	return
}
