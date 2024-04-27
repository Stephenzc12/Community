package logic

import (
	"user-management/dao/mysql"
	"user-management/model"
	"user-management/pkg/snowflake"
)

func CreateReply(r *model.Reply) (err error) {
	// 1.生成reply id
	r.RID = snowflake.GenID()
	// 2.保存到数据库
	return mysql.CreateReply(r)
	// 3.返回
}
