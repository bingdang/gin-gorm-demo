package middleware

import (
	"bookManager/dao"
	"bookManager/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		//从请求头中获取token
		ReqToken := c.Request.Header.Get("token")
		if len(ReqToken) == 0 {
			c.JSON(http.StatusForbidden, gin.H{
				"code": 90403,
				"msg":  "Request Token is empty",
				"data": nil,
			})
			c.Abort()
			return
		}
		//从库中校验
		userDAO := dao.NewUserDAO(dao.DB)
		auth := &model.User{
			Token: ReqToken,
		}
		err := userDAO.GetToken(auth)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"code": 90403,
				"msg":  "The token is incorrect",
				"data": nil,
			})
			c.Abort()
			return
		}
		c.Set("auth", true)
	}
}
