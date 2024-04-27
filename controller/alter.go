package controller

import (
	"github.com/gin-gonic/gin"
	"user-management/logic"
	"user-management/model"
)

func AlterHandler(c *gin.Context) {
	// 1.获取参数
	var a *model.Alter
	a = new(model.Alter)
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(404, gin.H{
			"code": 404,
			"msg":  "请求参数有误",
		})
	}

	// 2.修改数据库中的数据
	if err := logic.Alter(a); err != nil {
		c.JSON(404, gin.H{
			"code": 404,
			"msg":  "修改失败",
		})
	}
	// 3.返回响应
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "修改成功",
	})
}
