package dao

import (
	"log"

	"github.com/yzbtdiy/uni-update-server/models"
)

// 添加设备信息到 device_info 表中, 默认为非激活状态
func DbCacheDevice(device_id, device_mode string) {
	cacheDevice := models.DbDeviceInfoTable{
		DeviceId:   device_id,
		DeviceMode: device_mode,
	}
	result := db.Create(&cacheDevice)
	if result.RowsAffected == 1 {
		log.Println(cacheDevice.DeviceId + " add to db success, need to activate")
	}
}

// 验证设备, 返回两个布尔值, 第一个值为设备在数据库中是否存在记录, 存在为 true, 不存在为 false
// 第二个为值设备激活状态, 激活为 true, 否则为 false
func VerifyDevice(deviceId string) (bool, bool) {
	var deviceInfo *models.DbDeviceInfoTable
	result := db.Where("device_id = ?", deviceId).First(&deviceInfo)
	if result.RowsAffected == 1 && deviceInfo.Activate == 1 {
		return true, true
	} else if result.RowsAffected == 1 && deviceInfo.Activate == 0 {
		return true, false
	} else {
		return false, false
	}
}

// 读取 device_info 表中的数据
func GetDevListFromDb() (devList *[]models.DbDeviceInfoTable) {
	result := db.Find(&devList)
	if result.RowsAffected != 0 {
		return
	} else {
		return nil
	}
}

// 修改 device_info 表中的数据, 仅修改激活状态
func ChangeActivate(deviceId string, activateCode int) bool {
	var deviceInfo *models.DbDeviceInfoTable
	result := db.Model(&deviceInfo).Where("device_id = ?", deviceId).Update("activate", activateCode)
	if result.RowsAffected == 1 {
		return true
	} else {
		return false
	}
}

// 同步客户端版本号(已取消)
// func SyncEditionNumber(deviceId string, editionNumber int) bool {
// 	var deviceInfo *models.DbDeviceInfoTable
// 	result := db.Model(&deviceInfo).Where("device_id = ?", deviceId).Update("edition_number", editionNumber)
// 	if result.RowsAffected == 1 {
// 		return true
// 	} else {
// 		return false
// 	}
// }
