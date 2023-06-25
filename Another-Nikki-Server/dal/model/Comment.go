package model

import (
	"gorm.io/gorm"
)

const TableNameComment = "blog"

type Blog struct {
	gorm.Model
	Content string `json:"content"`
}

// TableName Comment's table name
func (*Blog) TableName() string {
	return TableNameComment
}
