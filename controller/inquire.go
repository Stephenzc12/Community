package controller

import (
	"github.com/gin-gonic/gin"
	"user-management/logic"
	"user-management/model"
)

func InquireHandler(c *gin.Context) {
	// 1.获取参数并获取参数
	var i *model.Inquiry
	i = new(model.Inquiry)

	if err := c.ShouldBindJSON(&i); err != nil {
		c.JSON(404, gin.H{
			"code": 404,
			"msg":  "请求参数有误",
		})
	}
	// 2.查询数据
	if err := logic.Inquire(i); err != nil {
		c.JSON(404, gin.H{
			"code": 404,
			"msg":  "查询失败",
		})
		return
	}
	// 3.返回响应
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "查询成功",
	})
}
