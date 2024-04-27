package mysql

import (
	"fmt"
	"user-management/model"
)

func CreateQuestion(p *model.Question) (err error) {
	sqlStr := `insert into question(
	question_id,author_id,question,create_time)
	values (?,?,?,?)`

	err = db.Exec(sqlStr, p.QID, p.AuthorID, p.Question, p.CreateTime).Error
	fmt.Printf("%+v\n", p)
	return
}

func CreateTableQuestion() (err error) {

	createTableQuestion := `CREATE TABLE IF NOT EXISTS question
	(
		question_id VARCHAR(255),
		author_id VARCHAR(255),
		question VARCHAR(255),
		create_time TIMESTAMP
	)`

	err = db.Exec(createTableQuestion).Error
	if err != nil {
		fmt.Println("drop table question error:", err)
		return err
	}
	return
}
