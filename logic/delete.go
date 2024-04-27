package logic

import (
	"github.com/gin-gonic/gin"
	"user-management/dao/mysql"
	"user-management/model"
)

func Delete(c *gin.Context, d *model.Delete) (err error) {
	// 在数据库中删除数据
	_, exists := c.Get("Q")
	if !exists {
		if err = mysql.DeleteReply(d); err != nil {
			return err
		}
	} else {
		if err = mysql.DeleteQuestion(d); err != nil {
			return err
		}
		if err = mysql.DeleteReply(d); err != nil {
			return err
		}
	}
	return
}
