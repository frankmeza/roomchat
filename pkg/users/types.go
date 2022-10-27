package users

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type (
	UserDB struct {
		ID        int            `gorm:"primaryKey;autoIncrement" json:"id"`
		CreatedAt time.Time      `json:"created_at"`
		UpdatedAt time.Time      `json:"updated_at"`
		DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	}

	User struct {
		UserDB
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
		return errors.New(
			fmt.Sprint("error on userProps.Scan", incomingValue),
		)
	}

	return json.Unmarshal([]byte(valueAsByteSlice), userProps)
}

func (userProps UserProps) Value() (driver.Value, error) {
	value, err := json.Marshal(&userProps)
	if err != nil {
		return nil, err
	}

	return value, nil
}
