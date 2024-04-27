package mysql

import "user-management/model"

func Alter(a *model.Alter) (err error) {
	sqlStr1 := "update question set question=? where question_id=?"
	if err = db.Exec(sqlStr1, a.NewQuestion, a.QID).Error; err != nil {
		return err
	}
	sqlStr2 := "update reply set reply=? where reply_id=?"
	if err = db.Exec(sqlStr2, a.NewReply, a.RID).Error; err != nil {
		return err
	}
	return
}
