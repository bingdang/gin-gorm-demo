package main

import (
	"bookManager/dao"
	"bookManager/router"
	"github.com/gin-gonic/gin"
)

func main() {
	//初始化数据库
	dao.InitMysql()

	//初始化gin
	GenEngine := gin.Default()

	//注册路由
	router.Router.InitApiRouter(GenEngine)

	//启动
	GenEngine.Run(":8000")
}
