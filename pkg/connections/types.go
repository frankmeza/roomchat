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
		ConnectionProps ConnectionProps `gorm:"type:jsonb" json:"connection_props"`
		Messages        []Message       `gorm:"type:jsonb" json:"messages"`
		Uuid            string          `json:"uuid"`
		db.DbRecord
	}

	ConnectionProps struct {
		FromUser   string `json:"from_user"`
		IsAnswered bool   `json:"is_answered"`
		Location   string `json:"location"`
		ToUser     string `json:"to_user"`
		Uuid       string `json:"uuid"`
	}

	Message struct {
		FromUser string `json:"from_user"`
		ToUser   string `json:"to_user"`
		Text     string `json:"text"`
		db.ChildRecord
	}
)

// from https://gorm.io/docs/data_types.html#Implements-Customized-Data-Type
func (connectionProps *ConnectionProps) Scan(incomingValue interface{}) error {
	valueAsByteSlice, ok := incomingValue.([]byte)
	if !ok {
		return errata.CreateError("ConnectionProps Scan", errors.New(""))
	}

	return json.Unmarshal([]byte(valueAsByteSlice), connectionProps)
}

func (connectionProps ConnectionProps) Value() (driver.Value, error) {
	value, err := json.Marshal(&connectionProps)
	if err != nil {
		return nil, errata.CreateError("ConnectionProps Value", err)
	}

	return value, nil
}

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
		return nil, errata.CreateError("Message Value", err)
	}

	return value, nil
}
