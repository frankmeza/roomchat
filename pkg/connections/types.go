package connections

import "github.com/frankmeza/roomchat/pkg/db"

type (
	Message struct {
		Connection string `json:"connection_id"`
		FromUser   string `json:"from_user"`
		ToUser     string `json:"to_user"`
		Text       string `json:"text"`
	}

	Connection struct {
		db.DbRecord
		FromUser   string    `json:"from_user"`
		IsAnswered bool      `json:"is_answered"`
		Location   string    `json:"location"`
		Messages   []Message `json:"messages"`
		ToUser     string    `json:"to_user"`
	}
)
