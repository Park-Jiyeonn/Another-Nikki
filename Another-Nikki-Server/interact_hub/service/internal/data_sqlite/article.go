package data_sqlite

import (
	"Another-Nikki/interact_hub/service/internal/biz"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
	"time"
)

type ArticleDataSqlite struct {
	ProblemId int64     `db:"problem_id"`
	CreatedAt time.Time `db:"created_time"`
	UpdatedAt time.Time `db:"updated_time"`
}

func InitArticlesTable(db *sqlx.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS articles (
   		article_id INTEGER PRIMARY KEY AUTOINCREMENT,
   		article_title TEXT NOT NULL,
   		article_description TEXT,
   		article_content TEXT,
   		created_time DATETIME DEFAULT CURRENT_TIMESTAMP,
   		updated_time DATETIME DEFAULT CURRENT_TIMESTAMP
   );`
	_, err := db.Exec(schema)
	return err
}

type articleImpl struct {
	db *sqlx.DB
}

func NewArticleRepo(data *Data) biz.ArticleRepo {
	if err := InitArticlesTable(data.GlobalDB); err != nil {
		panic(err)
	}
	return &articleImpl{
		db: data.GlobalDB,
	}
}

func (s *articleImpl) PostArticle(ctx context.Context, req *biz.PostArticleReq) (err error) {
	_, err = s.db.ExecContext(ctx,
		`INSERT INTO articles (article_title, article_description, article_content, updated_time)
		 VALUES (?, ?, ?, CURRENT_TIMESTAMP)`,
		req.ArticleTitle, req.ArticleDescription, req.ArticleContent)
	return
}

func (s *articleImpl) UpdateArticle(ctx context.Context, req *biz.UpdateArticleReq) (err error) {
	sqlStr := `UPDATE articles 
		SET article_title = ?, article_description = ?, article_content = ?, updated_time = CURRENT_TIMESTAMP 
		WHERE article_id = ?`
	_, err = s.db.ExecContext(ctx, sqlStr,
		req.ArticleTitle, req.ArticleDescription, req.ArticleContent, req.ArticleId)
	return
}

func (s *articleImpl) GetArticleById(ctx context.Context, req *biz.GetArticleByIdReq) (resp *biz.GetArticleByIdResp, err error) {
	resp = new(biz.GetArticleByIdResp)
	err = s.db.GetContext(ctx, resp, `
		SELECT article_title, article_description, article_content, created_time 
		FROM articles 
		WHERE article_id = ?`, req.ArticleId)
	return
}

func (s *articleImpl) GetArticleByPage(ctx context.Context, req *biz.GetArticleByPageReq) (resp *biz.GetArticleByPageResp, err error) {
	resp = new(biz.GetArticleByPageResp)
	rows, err := s.db.QueryxContext(ctx, `
		SELECT article_id, article_title, created_time, article_description 
		FROM articles 
		ORDER BY created_time DESC 
		LIMIT ? OFFSET ?`,
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
