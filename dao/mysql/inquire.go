package mysql

import "user-management/model"

func Inquire(i *model.Inquiry) (err error) {
	sqlStr1 := "create table if not exists auth_question as " +
		"select q.question_id, q.question, r.reply_id, r.reply " +
		"from question q " +
		"left join reply r on q.question_id = r.question_id " +
		"where q.author_id = ?"
	if err = db.Exec(sqlStr1, i.UserID).Error; err != nil {
		return err
	}
	sqlStr2 := "create table if not exists auth_reply as " +
		"select r.replier_id, r.reply_id, r.reply " +
		"from reply r " +
		"right join question q on r.question_id=q.question_id " +
		"where r.replier_id = ?"
	if err = db.Exec(sqlStr2, i.UserID).Error; err != nil {
		return err
	}
	return
}
