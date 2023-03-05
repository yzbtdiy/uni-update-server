package server

import (
	"net/http"

	"github.com/yzbtdiy/uni-update-server/api"
)

// 路由注册
func Router() {
	// App 客户端请求 api
	http.HandleFunc("/api/client_update", api.ClientUpdate)
	http.HandleFunc("/api/client_applist", api.ClientAppList)

	// 服务端请求 api
	http.HandleFunc("/api/login", api.Login)
	http.HandleFunc("/api/get_devlist", api.GetDevList)
	http.HandleFunc("/api/change_activate", api.ModDevInfo)
	http.HandleFunc("/api/get_version", api.GetVerInfo)
	http.HandleFunc("/api/add_version", api.AddVerInfo)
	http.HandleFunc("/api/mod_version", api.ModVerInfo)
	http.HandleFunc("/api/get_applist", api.GetAppList)
	http.HandleFunc("/api/add_applist", api.AddAppList)
	http.HandleFunc("/api/mod_applist", api.ModAppList)

	// web管理静态页面渲染
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./front-end/dist"))))
}
