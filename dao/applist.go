package dao

import (
	"github.com/yzbtdiy/uni-update-server/models"
)

// 读取 app_list 表中的数据
func GetAppListFromDb() (appList *[]models.DbAppListTable) {
	result := db.Find(&appList)
	if result.RowsAffected != 0 {
		return appList
	} else {
		return nil
	}
}

// 在 app_list 表中添加新数据
func AddAppInfoToDb(appInfo *models.DbAppListTable) bool {
	result := db.Create(&appInfo)
	if result.RowsAffected != 0 {
		return true
	} else {
		return false
	}
}

// 更新 app_list 表中的数据
func ModAppInfoToDb(appInfo *models.DbAppListTable) bool {
	// result := db.Model(models.DbAppListTable{}).Where("id = ?", appInfo.Id).Updates(models.AppInfo{Name: appInfo.Name, Url: appInfo.Url})
	result := db.Model(models.DbAppListTable{}).Where("id = ?", appInfo.Id).Updates(appInfo.AppInfo)
	if result.RowsAffected != 0 {
		return true
	} else {
		return false
	}
}
