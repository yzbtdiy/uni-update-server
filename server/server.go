package server

import (
	"log"
	"net/http"

	"github.com/yzbtdiy/uni-update-server/models"
	"github.com/yzbtdiy/uni-update-server/utils"
)

type UniUpdateServer struct{}

var App = &UniUpdateServer{}

// 启动 web 服务, 监听端口从配置文件读取
func (*UniUpdateServer) Start() {
	conf := utils.ReadYaml("./config/config.yaml", &models.Config{})
	server := http.Server{
		Addr: conf.Server.Ip + ":" + conf.Server.Port,
	}
	Router()
	log.Println("Server is runing on http://" + conf.Server.Ip + ":" + conf.Server.Port)
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
