package model

import "time"

type Question struct {
	QID        int64     `json:"question_id"`
	AuthorID   int64     `json:"author_id"`
	Question   string    `json:"question" binding:"required"`
	CreateTime time.Time `json:"create_time"`
}
