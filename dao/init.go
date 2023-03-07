package dao

import (
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"github.com/yzbtdiy/uni-update-server/models"
)

var db *gorm.DB

// 初始化数据库
func InitDb() {
	var err error
	db, err = gorm.Open(sqlite.Open("./config/app.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(1)

	tableInit(&models.DbDeviceInfoTable{})
	tableInit(&models.DbVersionInfoTable{})
	tableInit(&models.DbUserTable{})
	tableInit(&models.DbAppListTable{})
}

// 检查表是否存在, 不存在则创建
func tableInit(table interface{}) {
	if db.Migrator().HasTable(table) {
		return
	} else {
		db.Migrator().CreateTable(table)
	}
}
