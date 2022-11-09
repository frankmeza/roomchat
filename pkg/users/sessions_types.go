package users

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/frankmeza/roomchat/pkg/db"
	"github.com/frankmeza/roomchat/pkg/errata"
)

type (
	UserSession struct {
		UserSessionProps UserSessionProps `gorm:"type:jsonb" json:"user_session_props"`
		Uuid             string           `json:"uuid"`
		db.DbRecord
	}

	UserSessionProps struct {
		UserUuid string `json:"user_uuid"`
		Location string `json:"location"`
	}
)

func (userSessionProps *UserSessionProps) Scan(incomingValue interface{}) error {
	valueAsByteSlice, ok := incomingValue.([]byte)
	if !ok {
		return errata.CreateError(errors.New(""), errata.ErrMessage{
			Text: "UserSessionProps Scan",
		})
	}

	return json.Unmarshal([]byte(valueAsByteSlice), userSessionProps)
}

func (userSessionProps UserSessionProps) Value() (driver.Value, error) {
	value, err := json.Marshal(&userSessionProps)
	if err != nil {
		return nil, errata.CreateError(err, errata.ErrMessage{
			Text: "UserSessionProps Value",
		})
	}

	return value, nil
}
