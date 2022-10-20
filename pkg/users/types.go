package users

import (
	"gorm.io/datatypes"
)

type (
	User struct {
		ID        int `gorm:"primaryKey;autoIncrement"`
		UserSpecs datatypes.JSONMap
	}

	UserProps struct {
		Bio        string `json:"bio"`
		CurrentPic string `json:"current_pic"`
		Email      string `json:"email"`
		ID         string `json:"id"`
		Location   string `json:"location"`
		Name       string `json:"name"`
		Password   string `json:"-"`
		Pics       string `json:"pics"`
	}
)
