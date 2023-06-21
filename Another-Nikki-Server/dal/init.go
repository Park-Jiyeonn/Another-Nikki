package dal

import (
	"Another-Nikki/dal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	if DB != nil {
		return
	}
	dsn := "jiyeon:1234@tcp(127.0.0.1:3307)/Another_Nikki?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(
		&model.Blog{},
		&model.Log{},
	)
	if err != nil {
		panic(err)
	}
}
