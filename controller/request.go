package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"time"
	"user-management/middleware"
	"user-management/model"
)

func getCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(middleware.CtxUserIDKey)
	if !ok {
		return 0, errors.New("请先登录")
	}
	userID = uid.(int64)
	return
}

func getTimeNow1(p *model.Question) {
	p.CreateTime = time.Now()
}
func getTimeNow2(r *model.Reply) {
	r.CreateTime = time.Now()
}
