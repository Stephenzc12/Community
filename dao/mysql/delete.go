package mysql

import (
	"user-management/model"
)

func DeleteQuestion(d *model.Delete) (err error) {

	sqlStr1 := "delete from question where question_id = ?"
	sqlStr2 := "delete from reply where question_id = ?"
	err = db.Exec(sqlStr1, d.QID).Error
	if err != nil {
		return err
	}
	err = db.Exec(sqlStr2, d.QID).Error
	if err != nil {
		return err
	}
	return
}

func DeleteReply(d *model.Delete) error {

	sqlStr := "delete from reply where reply_id = ?"
	return db.Exec(sqlStr, d.RID).Error

}
