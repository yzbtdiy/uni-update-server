package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/yzbtdiy/uni-update-server/dao"
	"github.com/yzbtdiy/uni-update-server/models"
)

// 处理 /api/get_version 请求, 获取所有版本信息
func GetVerInfo(w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	isLogin := VerifyToken(r)
	if r.Method == "POST" {
		if isLogin {
			verInfos := dao.GetVerInfoFromDb()
			jsonData, _ := json.Marshal(verInfos)
			_, err := w.Write(jsonData)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

// 处理 /api/add_version 请求, 新增版本信息
func AddVerInfo(w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	if r.Method == "POST" {
		isLogin := VerifyToken(r)
		if isLogin {
			verInfo := GetPostJsonParams(r, &models.DbVersionInfoTable{})
			isAdd := dao.AddVerInfoToDb(verInfo)
			if isAdd {
				ReqSuccess(w, "add success")
			} else {
				ReqError(w, errors.New("add faild"))
			}
		}
	}
}

// 处理 /api/mod_version 请求, 修改版本信息
func ModVerInfo(w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	if r.Method == "POST" {
		isLogin := VerifyToken(r)
		if isLogin {
			params := GetPostJsonParams(r, &models.DbVersionInfoTable{})
			isMod := dao.ModVerInfoToDb(params)
			if isMod {
				ReqSuccess(w, "mod success")
			} else {
				ReqError(w, errors.New("mod faild"))
			}
		}
	}
}
