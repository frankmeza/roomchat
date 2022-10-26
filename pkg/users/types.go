package users

import "gorm.io/datatypes"

type (
	User struct {
		ID        int `gorm:"primaryKey;autoIncrement"`
		UserProps datatypes.JSONMap
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
	}
)
