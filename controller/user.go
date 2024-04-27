package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"user-management/logic"
	"user-management/model"
)

// SignUpHandler 处理注册请求的函数
func SignUpHandler(c *gin.Context) {

	// 1.获取参数和参数校验
	var p *model.ParamSignUp

	// ShouldBindJSON只能识别是否是JSON格式的
	if err := c.ShouldBindJSON(&p); err != nil {
		// 请求参数有误，直接返回响应
		c.JSON(404, gin.H{
			"code": 404,
			"msg":  "请求参数有误",
		})
		return
	}
	// 手动对参数进行一个校验
	if len(p.Username) == 0 {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "用户名不能为空",
		})
		return
	}
	if len(p.Password) < 6 {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "密码不能少于6位",
		})
		return
	}
	if len(p.RePassword) == 0 {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "请再次输入密码",
		})
		return
	}
	if p.RePassword != p.Password {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "两次密码输入不相同",
		})
		return
	}
	fmt.Println(p)

	// 2.业务处理
	if err := logic.SingUp(p); err != nil {
		c.JSON(404, gin.H{
			"code": 404,
			"msg":  "注册失败",
		})
		return
	}
	// 3.返回响应
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "注册成功",
	})
}

// LoginHandler 处理登录请求的函数
func LoginHandler(c *gin.Context) {

	// 1.参数获取和参数校验
	var p *model.ParamLogin

	if err := c.ShouldBind(&p); err != nil {
		c.JSON(404, gin.H{
			"code": 404,
			"msg":  "参数请求有误",
		})
		return
	}

	if len(p.Username) == 0 || len(p.Password) == 0 {
		c.JSON(404, gin.H{
			"code": 404,
			"msg":  "用户名和密码不能为空",
		})
		return
	}

	// 2.业务处理
	token, err := logic.Login(p)
	if err != nil {
		fmt.Println(err)
		c.JSON(404, gin.H{
			"code": 404,
			"msg":  "用户名或密码错误",
		})
		return
	}

	// 3.返回响应
	c.JSON(200, gin.H{
		"code":  200,
		"msg":   "登陆成功",
		"token": token,
	})
}
