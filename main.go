package main

import (
	"github.com/yzbtdiy/uni-update-server/dao"
	"github.com/yzbtdiy/uni-update-server/server"
	"github.com/yzbtdiy/uni-update-server/utils"
)

func init() {
	utils.InitConfig()
	dao.InitDb()
}

func main() {
	server.App.Start()
}
