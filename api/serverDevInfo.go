package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/yzbtdiy/uni-update-server/dao"
	"github.com/yzbtdiy/uni-update-server/models"
)

// 处理 /api/get_devlist 请求, 获取所有客户端设备信息
func GetDevList(w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	if r.Method == "POST" {
		isLogin := VerifyToken(r)
		if isLogin {
			devList := dao.GetDevListFromDb()
			jsonData, _ := json.Marshal(devList)
			_, err := w.Write(jsonData)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

// 处理 /api/mod_devlist 请求, 修改设备激活状态
func ModDevInfo(w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	if r.Method == "POST" {
		isLogin := VerifyToken(r)
		if isLogin {
			params := GetPostJsonParams(r, &models.PostModDevInfo{})
			devList := dao.ChangeActivate(params.DeviceId, params.Activate)
			jsonData, _ := json.Marshal(devList)
			_, err := w.Write(jsonData)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
