package model

import "gorm.io/gorm"

type Problem struct {
	gorm.Model
	Name    string `json:"name"`
	Content string `json:"content"`
}

func (*Problem) TableName() string {
	return "problem"
}
