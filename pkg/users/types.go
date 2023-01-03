package users

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/frankmeza/roomchat/pkg/db"
	"github.com/frankmeza/roomchat/pkg/errata"
	"github.com/frankmeza/roomchat/pkg/sessions"
)

type (
	User struct {
		UserProps UserProps `gorm:"type:jsonb" json:"user_props"`
		Uuid      string    `json:"uuid"`
		db.DbRecord
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

type (
	GetUserParams struct {
		Email     string
		ID        string
		ParamName string
		Username  string
		Uuid      string
	}

	handleLoginMacroMetadata struct {
		session sessions.Session
		token   string
	}

	handleLoginParams struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Username string `json:"username"`
	}

	UsersAPI struct {
		apiType string
	}
)

// from https://gorm.io/docs/data_types.html#Implements-Customized-Data-Type
func (userProps *UserProps) Scan(incomingValue interface{}) error {
	valueAsByteSlice, isOk := incomingValue.([]byte)
	if !isOk {
		return errata.CreateError(errors.New(""), []string{
			"UserProps Scan",
		})
	}

	return json.Unmarshal([]byte(valueAsByteSlice), userProps)
}

func (userProps UserProps) Value() (driver.Value, error) {
	value, err := json.Marshal(&userProps)
	if err != nil {
		return nil, errata.CreateError(err, []string{
			"UserProps Value",
		})
	}

	return value, nil
}
