package router

import (
	"github.com/gin-gonic/gin"
	"user-management/controller"
	"user-management/middleware"
)

func SetUpRoute() *gin.Engine {

	r := gin.Default()

	// 注册业务路由
	r.POST("/signup", controller.SignUpHandler)
	// 登陆业务路由
	r.POST("/login", controller.LoginHandler)
	// 发起问题业务路由
	r.POST("/question", middleware.JwtMiddleware(), controller.CreateQuestionHandler)
	// 用户回答业务路由
	r.POST("/reply", middleware.JwtMiddleware(), controller.CreateReplyHandler)
	// 用户删除业务路由
	r.POST("delete", middleware.JwtMiddleware(), controller.DeleteHandler)
	// 用户修改业务路由
	r.POST("alter", middleware.JwtMiddleware(), controller.AlterHandler)
	// 查询问题和回答路由
	r.POST("/inquire", middleware.JwtMiddleware(), controller.InquireHandler)

	return r
}
