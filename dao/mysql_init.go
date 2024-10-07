package dao

import (
	"bookManager/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

// 初始化数据库
func InitMysql() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/books?charset=utf8mb4&parseTime=True&loc=Local"
	mdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败" + err.Error())
	}
	DB = mdb
	log.Printf("连接数据库成功%s", dsn)

	//自动创建表
	if err := DB.AutoMigrate(model.User{}, model.Book{}); err != nil {
		log.Printf("自动创建表失败%s", err.Error())
	}

}
