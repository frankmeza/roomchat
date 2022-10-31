package users

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/frankmeza/roomchat/pkg/db"
	"github.com/frankmeza/roomchat/pkg/errata"
)

type (
	User struct {
		db.DbRecord
		UserProps UserProps `gorm:"type:jsonb" json:"user_props"`
		Uuid      string    `json:"uuid"`
	}

	UserProps struct {
		Bio        string `json:"bio"`
		CurrentPic string `json:"current_pic"`
		Email      string `json:"email"`
		Location   string `json:"location"`
		Name       string `json:"name"`
		Password   string `json:"password"`
		Pics       string `json:"pics"`
		Username   string `json:"username"`
		Uuid       string `json:"uuid"`
	}
)

// from https://gorm.io/docs/data_types.html#Implements-Customized-Data-Type
func (userProps *UserProps) Scan(incomingValue interface{}) error {
	valueAsByteSlice, ok := incomingValue.([]byte)
	if !ok {
		return errata.CreateError(errata.ErrataParams{
			Err:    errors.New(fmt.Sprint("", incomingValue)),
			ErrMsg: "UserProps Scan",
		})
	}

	return json.Unmarshal([]byte(valueAsByteSlice), userProps)
}

func (userProps UserProps) Value() (driver.Value, error) {
	value, err := json.Marshal(&userProps)
	if err != nil {
		return nil, errata.CreateError(errata.ErrataParams{
			Err:    err,
			ErrMsg: "UserProps Scan",
		})
	}

	return value, nil
}
