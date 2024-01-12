package model

import (
	"gorm.io/gorm"
	"time"
)

type Judge struct {
	ID            int64 `gorm:"primarykey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	UserID        int64          `json:"user_id"`
	UserName      string         `json:"user_name"`
	ProblemID     int64          `json:"problem_id"`
	ProblemName   string         `json:"problem_name"`
	Language      string         `json:"language"`
	Code          string         `json:"code"`
	CompileStatus string         `json:"compile_status"`
	JudgeStatus   string         `json:"judge_status"`
}

func (*Judge) TableName() string {
	return "judge"
}
