package connections

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/frankmeza/roomchat/pkg/db"
	"github.com/frankmeza/roomchat/pkg/errata"
)

type (
	Connection struct {
		FromUser   string    `json:"from_user"`
		IsAnswered bool      `json:"is_answered"`
		Location   string    `json:"location"`
		Messages   []Message `gorm:"type:jsonb" json:"messages"`
		ToUser     string    `json:"to_user"`
		Uuid       string
		db.DbRecord
	}

	Message struct {
		FromUser string `json:"from_user"`
		ToUser   string `json:"to_user"`
		Text     string `json:"text"`
		db.ChildRecord
	}
)

// from https://gorm.io/docs/data_types.html#Implements-Customized-Data-Type
func (message *Message) Scan(incomingValue interface{}) error {
	valueAsByteSlice, ok := incomingValue.([]byte)
	if !ok {
		return errata.CreateError("Message Scan", errors.New(""))
	}

	return json.Unmarshal([]byte(valueAsByteSlice), message)
}

func (message Message) Value() (driver.Value, error) {
	value, err := json.Marshal(&message)
	if err != nil {
		return nil, errata.CreateError("Message Scan", err)
	}

	return value, nil
}
