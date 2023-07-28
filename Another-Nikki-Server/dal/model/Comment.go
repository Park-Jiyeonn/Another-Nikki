package model

import (
	"gorm.io/gorm"
)

const TableNameComment = "comment"

type Comment struct {
	gorm.Model
	Content      string    `json:"content"`
	ArticleID    int       `json:"article_id" gorm:"index:index_article_id"`
	AuthorID     int       `json:"author_id"`
	AuthorAvatar string    `json:"author_avatar"`
	AuthorName   string    `json:"author_name"`
	RootID       int       `json:"root_id" gorm:"index:root_id"`
	ParentID     int       `json:"parent_id"`
	ParentName   string    `json:"parent_name"`
	Children     []Comment `gorm:"-" json:"children"`
}

// TableName Comment's table name
func (*Comment) TableName() string {
	return TableNameComment
}
