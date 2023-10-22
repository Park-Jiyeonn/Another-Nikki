package model

import (
	"gorm.io/gorm"
	"time"
)

const TableNameUser = "user"

type User struct {
	gorm.Model
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	Avatar        string    `json:"avtar"`
	VisitTime     int64     `json:"visit_time"`
	LastVisitTime time.Time `json:"last_visit_time"`
}

// TableName Comment's table name
func (*User) TableName() string {
	return TableNameUser
}
