package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/yzbtdiy/uni-update-server/dao"
	"github.com/yzbtdiy/uni-update-server/models"
)

// 处理 /api/client_applist 请求, 返回软件列表供客户端渲染
func ClientAppList(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		params := GetPostJsonParams(r, &models.PostAppList{})
		hasDevice, isActivate := dao.VerifyDevice(params.DeviceId)
		if hasDevice {
			if isActivate {
				appList := dao.GetAppListFromDb()
				w.Header().Set("Content-Type", "application/json")
				jsonData, _ := json.Marshal(appList)
				_, err := w.Write(jsonData)
				if err != nil {
					log.Println(err)
				}
			} else {
				log.Println("Find device " + params.DeviceId + " in db, but not activate")
			}
		} else {
			log.Println("No such device, please add it")
		}
	}
}
