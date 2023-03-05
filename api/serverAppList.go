package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/yzbtdiy/uni-update-server/dao"
	"github.com/yzbtdiy/uni-update-server/models"
)

// 处理 /api/get_applist 请求, 获取软件列表
func GetAppList(w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	if r.Method == "POST" {
		isLogin := VerifyToken(r)
		if isLogin {
			appList := dao.GetAppListFromDb()
			jsonData, _ := json.Marshal(appList)
			_, err := w.Write(jsonData)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

// 处理 /api/add_applist 请求, 添加软件列表
func AddAppList(w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	if r.Method == "POST" {
		isLogin := VerifyToken(r)
		if isLogin {
			params := GetPostJsonParams(r, &models.DbAppListTable{})
			isAdd := dao.AddAppInfoToDb(params)
			if isAdd {
				ReqSuccess(w, "Add success")
			} else {
				ReqError(w, errors.New("add faild"))
			}
		}
	}
}

// 处理 /api/mod_applist 请求, 修改软件列表
func ModAppList(w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	if r.Method == "POST" {
		isLogin := VerifyToken(r)
		if isLogin {
			params := GetPostJsonParams(r, &models.DbAppListTable{})
			isMod := dao.ModAppInfoToDb(params)
			if isMod {
				ReqSuccess(w, "mod success")
			} else {
				ReqError(w, errors.New("mod faild"))
			}
		}
	}
}
