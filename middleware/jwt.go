package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
	"user-management/pkg/jwt"
)

const CtxUserIDKey = "userID"

func JwtMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URL
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(404, gin.H{
				"code": 404,
				"msg":  "请求头中auth为空",
			})
			c.Abort()
			return
		}
		// 按空格分割 判断格式是否有误
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 && parts[0] != "Bearer" {
			c.JSON(404, gin.H{
				"code": 404,
				"msg":  "请求头中auth格式有误",
			})
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，使用之前定义好的解析JWT的函数来解析
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			c.JSON(404, gin.H{
				"code": 404,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}

		c.Set(CtxUserIDKey, mc.UserID)

		c.Next()
	}

}
