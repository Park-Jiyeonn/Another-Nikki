package model

import (
	"gorm.io/gorm"
)

const TableNameComment = "comment"

type Comment struct {
	gorm.Model
	Content      string    `json:"content"`
	AuthorID     int       `json:"author_id"`
	AuthorAvatar string    `json:"author_avatar"`
	AuthorName   string    `json:"author_name"`
	RootID       int       `json:"root_id"`
	ParentID     int       `json:"parent_id"`
	ParentName   string    `json:"parent_name"`
	Children     []Comment `gorm:"-" json:"children"`
}

// TableName Comment's table name
func (*Comment) TableName() string {
	return TableNameComment
}
