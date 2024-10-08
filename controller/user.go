package controller

import (
	"fmt"
	"gin-gorm-demo/dao"
	"gin-gorm-demo/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 注册模块
func RegisterHandler(c *gin.Context) {
	userDAO := dao.NewUserDAO(dao.DB)
	p := new(model.User)
	//context type: application/json
	if err := c.ShouldBind(p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "错误的传参:" + err.Error(),
			"data": nil,
			"code": "90400",
		})
		return
	}

	//账号是否存在
	if err := userDAO.Has(p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "注册失败:" + err.Error(),
			"data": nil,
			"code": "90400",
		})
		return
	}

	//账号密码入库
	if err := userDAO.Add(p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "注册失败:" + err.Error(),
			"data": nil,
			"code": "90400",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "注册成功",
		"data": fmt.Sprintf("%s 注册成功，请牢记您的密码", p.Username),
		"code": "90200",
	})
}

// 登陆模块
func LoginHandler(c *gin.Context) {
	userDAO := dao.NewUserDAO(dao.DB)
	p := new(model.User)
	//context type: application/json
	if err := c.ShouldBind(p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "错误的传参:" + err.Error(),
			"data": nil,
			"code": "90400",
		})
		return
	}
	//验证账号密码
	u := &model.User{
		Username: p.Username,
		Password: p.Password,
	}

	//为空报错
	if rows := userDAO.VerifyAccountPassword(u); rows != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"msg":  "登陆失败，请检查账号密码",
			"data": nil,
			"code": "90403",
		})
		return
	}

	//生成token。用数据库存token做校验
	tx, tk := userDAO.UpdateToken(u)
	if tx != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "登陆失败，请联系管理员",
			"data": nil,
			"code": "90500",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":   "登陆成功",
		"token": tk,
	})
}
