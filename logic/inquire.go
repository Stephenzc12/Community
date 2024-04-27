package logic

import (
	"user-management/dao/mysql"
	"user-management/model"
)

func Inquire(i *model.Inquiry) (err error) {
	// 在数据库中查询用户数据
	if err = mysql.Inquire(i); err != nil {
		return err
	}
	return
}
