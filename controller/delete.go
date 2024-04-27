package controller

import (
	"github.com/gin-gonic/gin"
	"user-management/logic"
	"user-management/model"
)

func DeleteHandler(c *gin.Context) {
	// 1.获取参数并进行参数校验
	var d *model.Delete
	d = new(model.Delete)

	if err := c.ShouldBindJSON(&d); err != nil {
		c.JSON(404, gin.H{
			"code": 404,
			"msg":  "请求参数错误",
		})
		return
	}
	c.Set("Q", d.QID)
	c.Set("R", d.RID)

	// 2.删除记录
	if err := logic.Delete(c, d); err != nil {
		c.JSON(404, gin.H{
			"code": 404,
			"msg":  "删除失败",
		})
	}

	// 3.返回响应
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}
