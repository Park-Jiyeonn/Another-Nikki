package model

import (
	"gorm.io/gorm"
)

const TableNameUser = "user"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Avatar   string `json:"avtar"`
}

// TableName Comment's table name
func (*User) TableName() string {
	return TableNameUser
}
