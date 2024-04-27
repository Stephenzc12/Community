package logic

import (
	"user-management/dao/mysql"
	"user-management/model"
)

func Alter(a *model.Alter) (err error) {
	// 在数据库中修改数据
	if err = mysql.Alter(a); err != nil {
		return err
	}
	return
}
