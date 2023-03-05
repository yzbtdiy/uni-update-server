package api

import (
	"errors"
	"net/http"

	"github.com/yzbtdiy/uni-update-server/dao"
	"github.com/yzbtdiy/uni-update-server/models"
	"github.com/yzbtdiy/uni-update-server/utils"
)

// 处理 /api/login 请求, 登录管理界面
func Login(w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	if r.Method == "POST" {
		Params := GetPostJsonParams(r, &models.PostUser{})
		loginRes, err := TryLogin(Params.Username, Params.Password)
		if err != nil {
			ReqError(w, err)
			return
		}
		ReqSuccess(w, loginRes)
	}
}

// 尝试登录, 用户名密码正确则返回 token
func TryLogin(username, password string) (*models.ResLogin, error) {
	user := dao.GetUserInfoFromDB(username, password)
	if user != nil {
		token, err := utils.GenerateToken(user.Id, user.Username)
		if err != nil {
			return nil, errors.New("token生成失败")
		}
		resLogin := &models.ResLogin{
			Token: token,
		}
		return resLogin, nil
	} else {
		return nil, errors.New("用户不存在")
	}
}
