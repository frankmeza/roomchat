package sessions

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/frankmeza/roomchat/pkg/db"
	"github.com/frankmeza/roomchat/pkg/errata"
)

type (
	Session struct {
		SessionProps SessionProps `gorm:"type:jsonb" json:"user_session_props"`
		Uuid         string       `json:"uuid"`
		db.DbRecord
	}

	SessionProps struct {
		UserUuid string `json:"user_uuid"`
		Location string `json:"location"`
	}
)

type (
	SessionsAPI struct {
		apiType string
	}

	CreateSessionParams struct {
		Location string
		Uuid     string
	}
)

func (SessionProps *SessionProps) Scan(incomingValue interface{}) error {
	valueAsByteSlice, isOk := incomingValue.([]byte)
	if !isOk {
		return errata.CreateError(errors.New(""), []string{
			"SessionProps Scan",
		})
	}

	return json.Unmarshal([]byte(valueAsByteSlice), SessionProps)
}

func (SessionProps SessionProps) Value() (driver.Value, error) {
	value, err := json.Marshal(&SessionProps)
	if err != nil {
		return nil, errata.CreateError(err, []string{
			"SessionProps Value",
		})
	}

	return value, nil
}
