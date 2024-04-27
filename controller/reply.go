package controller

import (
	"github.com/gin-gonic/gin"
	"user-management/logic"
	"user-management/model"
)

func CreateReplyHandler(c *gin.Context) {
	// 1.获取参数及参数校验

	var r *model.Reply
	r = new(model.Reply)

	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(404, gin.H{
			"code": 404,
			"msg":  "请求参数有误",
		})
		return
	}

	// 从c里面获取当前发请求的用户ID
	var err error
	if r.ReplierID, err = getCurrentUserID(c); err != nil {
		c.JSON(404, gin.H{
			"code": 404,
			"msg":  "请先登录",
		})
		return
	}

	// 获取当前时间
	getTimeNow2(r)

	// 2.回答问题
	err = logic.CreateReply(r)
	if err != nil {
		c.JSON(404, gin.H{
			"code": 404,
			"msg":  "发表失败",
		})
		return
	}

	// 3.返回响应
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "发表成功",
	})
}
