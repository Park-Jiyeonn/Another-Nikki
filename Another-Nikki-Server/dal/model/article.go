package model

import "gorm.io/gorm"

type Article struct {
    gorm.Model
    Title string `json:"title"`
    Description string `json:"description"`
    Content string `json:"content"`
}

func (*Article) TableName() string{
    return "article"
}