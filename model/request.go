package model

type Delete struct {
	QID int64 `json:"question_id"`
	RID int64 `json:"reply_id"`
}

type Alter struct {
	QID         int64  `json:"question_id"`
	RID         int64  `json:"reply_id"`
	NewQuestion string `json:"new_question"`
	NewReply    string `json:"new_reply"`
}

type Inquiry struct {
	UserID   int64  `json:"user_id" binding:"required"`
	QID      int64  `json:"question_id"`
	RID      int64  `json:"reply_id"`
	Question string `json:"question"`
	Reply    string `json:"reply"`
}
