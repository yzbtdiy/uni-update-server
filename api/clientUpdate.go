package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/yzbtdiy/uni-update-server/dao"
	"github.com/yzbtdiy/uni-update-server/models"
	"github.com/yzbtdiy/uni-update-server/utils"
)

// 处理 /api/client_update 请求, 返回数据由客户端判断是否更新
func ClientUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		params := GetPostJsonParams(r, &models.PostDevice{})
		hasDevice, isActivate := dao.VerifyDevice(params.DeviceId)
		if hasDevice {
			if isActivate {
				updateId := utils.ReadYaml("./config/config.yaml", &models.Config{}).Client.UpdateId
				w.Header().Set("Content-Type", "application/json")
				versionData := dao.ClientGetUpdate(updateId).VersionInfo
				jsonData, _ := json.Marshal(versionData)
				_, err := w.Write(jsonData)
				if err != nil {
					log.Println(err)
				}

			} else {
				log.Println("Find device " + params.DeviceId + " in db, but not activate")
			}
		} else {
			log.Println("No such device, will add it but not activate")
			dao.DbCacheDevice(params.DeviceId, params.DeviceMode)
		}
	}
}
