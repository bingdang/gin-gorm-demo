package controller

import (
	"fmt"
	"gin-gorm-demo/dao"
	"gin-gorm-demo/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 新增书籍
func Add(c *gin.Context) {
	bookDAO := dao.NewBookDAO(dao.DB)
	p := new(model.Book)
	//context type: application/json
	if err := c.ShouldBind(p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "错误的传参:" + err.Error(),
			"data": nil,
			"code": "90400",
		})
		return
	}

	//书籍是否存在
	_, b, err := bookDAO.Has(p)
	if err != nil && b == true {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "添加书籍失败:" + err.Error(),
			"data": nil,
			"code": "90400",
		})
		return
	}

	//书籍入库
	if err := bookDAO.Add(p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "添加书籍失败:" + err.Error(),
			"data": nil,
			"code": "90400",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "添加书籍成功",
		"data": fmt.Sprintf("%s 书籍添加成功", p.Name),
		"code": "90200",
	})
}

// 书籍列表
func List(c *gin.Context) {
	bookDAO := dao.NewBookDAO(dao.DB)
	list, err := bookDAO.List()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "查询书籍列表失败:" + err.Error(),
			"data": nil,
			"code": "90400",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "查询书籍列表成功",
		"data": list,
		"code": "90200",
	})
}

// 获取单个书籍详情
func Get(c *gin.Context) {
	bookDAO := dao.NewBookDAO(dao.DB)
	p := new(model.Book)
	//Get请求
	if err := c.ShouldBind(p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "错误的传参:" + err.Error(),
			"data": nil,
			"code": "90400",
		})
		return
	}

	//获取书籍详情
	get, _, err := bookDAO.Get(p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "查询书籍详情失败:" + err.Error(),
			"data": nil,
			"code": "90400",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "查询书籍详情成功",
		"data": get,
		"code": "90200",
	})
}

// 更新书籍
func Update(c *gin.Context) {
	bookDAO := dao.NewBookDAO(dao.DB)
	p := new(model.Book)
	//context type: application/json
	if err := c.ShouldBind(p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "错误的传参:" + err.Error(),
			"data": nil,
			"code": "90400",
		})
		return
	}

	//根据id查单个
	_, _, err := bookDAO.Get(p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "查询书籍详情失败:" + err.Error(),
			"data": nil,
			"code": "90400",
		})
		return
	}

	//书籍更新
	if err := bookDAO.Update(p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "更新书籍失败:" + err.Error(),
			"data": nil,
			"code": "90400",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "更新书籍成功",
		"data": fmt.Sprintf("%s 书籍更新成功，", p.Name),
		"code": "90200",
	})
}

// 删除书籍
func Delete(c *gin.Context) {
	bookDAO := dao.NewBookDAO(dao.DB)
	p := new(model.Book)
	//context type: application/json
	if err := c.ShouldBind(p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "错误的传参:" + err.Error(),
			"data": nil,
			"code": "90400",
		})
		return
	}

	//根据id查单个
	_, _, err := bookDAO.Get(p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "查询书籍详情失败:" + err.Error(),
			"data": nil,
			"code": "90400",
		})
		return
	}

	//书籍删除
	if err := bookDAO.Delete(p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "删除书籍失败:" + err.Error(),
			"data": nil,
			"code": "90400",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "删除书籍成功",
		"data": fmt.Sprintf("%s 书籍删除成功，", p.Name),
		"code": "90200",
	})
}
