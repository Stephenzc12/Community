package logic

import (
	"user-management/dao/mysql"
	"user-management/model"
	"user-management/pkg/snowflake"
)

func CreateQuestion(p *model.Question) (err error) {
	// 1.生成question id和获取user id
	p.QID = snowflake.GenID()
	// 2.保存到数据库
	return mysql.CreateQuestion(p)
	// 3.返回

}
