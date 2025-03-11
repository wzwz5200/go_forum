package initdb

import (
	"fmt"
	"web/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ReDB *gorm.DB

func Initdb() *gorm.DB {

	fmt.Println("HELLOW")
	dsn := "host=localhost user=postgres password=18326307873qq dbname=test port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {

		fmt.Println("数据库连接错误", err.Error())

	}

	db.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{}, &model.Section{})

	ReDB = db

	return ReDB

}
