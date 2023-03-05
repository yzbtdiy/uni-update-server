package dao

import (
	"log"
	"strconv"

	"github.com/yzbtdiy/uni-update-server/models"
)

// 读取 version_info 表中的数据
func GetVerInfoFromDb() (versionInfo *[]models.DbVersionInfoTable) {
	result := db.Find(&versionInfo)
	if result.RowsAffected != 0 {
		return versionInfo
	} else {
		return nil
	}
}

// 在 version_info 表中添加数据
func AddVerInfoToDb(versionInfo *models.DbVersionInfoTable) bool {
	result := db.Create(&versionInfo)
	if result.RowsAffected != 0 {
		log.Println("版本号 " + strconv.Itoa(versionInfo.EditionNumber) + " 添加成功")
		return true
	} else {
		return false
	}
}

// 修改 version_info 表中数据
func ModVerInfoToDb(versionInfo *models.DbVersionInfoTable) bool {
	result := db.Model(models.DbVersionInfoTable{}).Where("id = ?", versionInfo.Id).Updates(versionInfo.VersionInfo)
	if result.RowsAffected != 0 {
		return true
	} else {
		return false
	}
}

func ClientGetUpdate(verId int) (verInfo *models.DbVersionInfoTable) {
	result := db.Where("id = ?", verId).First(&verInfo)
	if result.RowsAffected == 1 {
		return verInfo
	} else {
		return nil
	}
}
