package users

import (
	"gorm.io/datatypes"
)

type (
	User struct {
		ID        int `gorm:"primaryKey;autoIncrement"`
		UserProps datatypes.JSONMap
	}

	UserProps struct {
		Bio        string `gorm:"type:text" json:"bio"`
		CurrentPic string `gorm:"type:text" json:"current_pic"`
		Email      string `gorm:"type:text" json:"email"`
		Location   string `gorm:"type:text" json:"location"`
		Name       string `gorm:"type:text" json:"name"`
		Password   string `gorm:"type:text" json:"password"`
		Pics       string `gorm:"type:text" json:"pics"`
		Username   string `gorm:"type:text" json:"username"`
	}
)
