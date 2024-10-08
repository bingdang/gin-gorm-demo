package router

import (
	"gin-gorm-demo/controller"
	"gin-gorm-demo/middleware"

	"github.com/gin-gonic/gin"
)

var Router router

type router struct{}

func (*router) InitApiRouter(r *gin.Engine) {
	r.POST("/register", controller.RegisterHandler)
	r.POST("/login", controller.LoginHandler)
	bookGroup := r.Group("/book")
	{
		//注册token验证
		bookGroup.Use(middleware.AuthMiddleWare())
		//注册其他路由
		bookGroup.POST("/add", controller.Add)
		bookGroup.GET("/list", controller.List)
		bookGroup.GET("/detail", controller.Get)
		bookGroup.PUT("/update", controller.Update)
		bookGroup.DELETE("/delete", controller.Delete)
	}
}
