package users

import (
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
