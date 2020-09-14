package model

import (
	"dcgin/config"
	_ "github.com/go-sql-driver/mysql" // mysql驱动
	"github.com/jinzhu/gorm"
	"log"
)

// DB 实例
var DB *gorm.DB

func InitDb() {
	dbConnStr := config.AppConfig.MySqlDsn
	db, err := gorm.Open("mysql", dbConnStr)
	if err != nil {
		log.Panic("failed to connect database,err:" + err.Error())
	}
	db.DB().SetMaxOpenConns(200)

	db.SingularTable(true)
	if config.AppConfig.Env != "online" {
		db.LogMode(true)
	}
	db = db.Debug()

	DB = db
}

//GetDB 获取数据库连接对象
func GetDB() *gorm.DB {
	return DB
}

// Migrate 自动建表
func Migrate() {
	//DB.AutoMigrate(Ask{})
	//DB.AutoMigrate(Reply{})
}
