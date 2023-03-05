package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/yzbtdiy/uni-update-server/models"
	"github.com/yzbtdiy/uni-update-server/utils"
)

// 解析 post 请求参数
func GetPostJsonParams[T models.PostJson](r *http.Request, params T) T {
	body, _ := io.ReadAll(r.Body)
	_ = json.Unmarshal(body, &params)
	return params
}

// 验证 token 有效性
func VerifyToken(r *http.Request) bool {
	token := r.Header.Get("Authorization")
	// _, claim, err := utils.ParseToken(token)
	// id := claim.Id
	_, err := utils.ParseToken(token)
	if err != nil {
		// fmt.Println("token 无效")
		fmt.Println(err)
		return false
	} else {
		// fmt.Println("token 有效")
		return true
	}
}

// 处理跨域
func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Request-Method", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Lengsth, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")
}

// 请求成功返回
func ReqSuccess(w http.ResponseWriter, data interface{}) {
	var result models.ResData
	result.Code = 200
	result.Error = ""
	result.Data = data
	resultJson, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(resultJson)
	if err != nil {
		log.Println(err)
	}
}

// 请求失败返回
func ReqError(w http.ResponseWriter, err error) {
	var result models.ResData
	result.Code = 500
	result.Error = err.Error()
	resultJson, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resultJson)
	if err != nil {
		log.Println(err)
	}
}
