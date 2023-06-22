package model

import (
	"gorm.io/gorm"
)

const TableNameLog = "log"

type Log struct {
	gorm.Model
	Api 		string `json:"api"`
	Status 		int    `json:"status"`
	IP  		string `json:"ip"`
	Response 	string `json:"response"`
}

func (*Log) TableName() string {
	return TableNameLog
}
