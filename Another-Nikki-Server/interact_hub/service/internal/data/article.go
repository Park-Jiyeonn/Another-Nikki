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

type articleRepo struct {
	db *sqlx.DB
}

func NewArticleRepo(data *Data) biz.ArticleRepo {
	return &articleRepo{
		db: data.GlobalDB,
	}
}

// CREATE TABLE judge (
//    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
//    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
//    user_id BIGINT NOT NULL DEFAULT 0,
//    user_name VARCHAR(255) NOT NULL DEFAULT '',
//    problem_id BIGINT NOT NULL DEFAULT 0,
//    problem_name VARCHAR(255) NOT NULL DEFAULT '',
//    language VARCHAR(255) NOT NULL DEFAULT '',
//    code TEXT NOT NULL,
//    compile_status VARCHAR(255) NOT NULL DEFAULT 'in queue',
//    compile_log VARCHAR(255) NOT NULL DEFAULT '',
//    judge_status VARCHAR(255) NOT NULL DEFAULT '-',
//    cpu_time_used VARCHAR(255) NOT NULL DEFAULT '0',
//    memory_used VARCHAR(255) NOT NULL DEFAULT '0'
// );

func (s *articleRepo) CreateArticle(ctx context.Context, articleId int64, title, description, content string) (err error) {
	return
}
func (s *articleRepo) GetArticleById(ctx context.Context, articleId int64) (articleTitle, articleDescription, articleContent, createTime string, err error) {
	return
}
func (s *articleRepo) GetArticleByPage(ctx context.Context, pageNum, pageSize int64) (articles []*biz.Articles, err error) {
	return
}
