package model

import "time"

type Reply struct {
	QID        int64     `json:"question_id" binding:"required"`
	RID        int64     `json:"reply_id"`
	ReplierID  int64     `json:"replier_id"`
	Reply      string    `json:"reply" binding:"required"`
	CreateTime time.Time `json:"create_time"`
}
