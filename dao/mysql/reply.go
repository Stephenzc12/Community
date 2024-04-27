package mysql

import (
	"fmt"
	"user-management/model"
)

func CreateReply(r *model.Reply) (err error) {
	sqlStr1 := `insert into reply(
	question_id,replier_id,reply_id,reply,create_time)
	values(?,?,?,?,?)`

	err = db.Exec(sqlStr1, r.QID, r.ReplierID, r.RID, r.Reply, r.CreateTime).Error
	return
}

func CreateTableReply() (err error) {
	createTableReply := `CREATE TABLE IF NOT EXISTS reply
	(
		question_id VARCHAR(255),
		replier_id VARCHAR(255),
		reply_id VARCHAR(255),
		reply VARCHAR(255),
		create_time TIMESTAMP
	)`
	err = db.Exec(createTableReply).Error
	if err != nil {
		fmt.Println("create table reply err:", err)
		return err
	}
	return
}
