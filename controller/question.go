package controller

import (
	"github.com/gin-gonic/gin"
	"user-management/logic"
	"user-management/model"
)

func CreateQuestionHandler(c *gin.Context) {
	// 1.获取参数及参数校验
	var p *model.Question
	p = new(model.Question)

	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(404, gin.H{
			"code": 404,
			"msg":  "请求参数有误",
		})
		return
	}
	//从c里面获取当前发请求的用户ID
	userID, err := getCurrentUserID(c)
	if err != nil {
		c.JSON(404, gin.H{
			"code": 404,
			"msg":  "请先登录",
		})
		return
	}
	p.AuthorID = userID
	// 获取当前时间
	getTimeNow1(p)

	// 2.发起问题
	if err := logic.CreateQuestion(p); err != nil {
		c.JSON(404, gin.H{
			"code": 404,
			"msg":  "发起问题失败",
		})
		return
	}

	// 3.返回响应
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "发表成功",
	})
}
