package users

import (
	"errors"

	cc "github.com/frankmeza/roomchat/pkg/constants"
	"github.com/labstack/echo/v4"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type GetUserParams struct {
	Email string
	ID    string
}

func getUserByParam(
	conn *gorm.DB,
	user *User,
	paramsName string,
	params GetUserParams,
) (*User, error) {
	var paramToUse interface{}
	if paramsName == cc.EMAIL {
		paramToUse = params.Email
	}

	if paramsName == cc.ID {
		paramToUse = params.ID
	}

	query := datatypes.
		JSONQuery(cc.USER_SPECS).
		Equals(paramToUse, paramsName)

	result := conn.Debug().Find(&user, query)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &User{}, result.Error
	}

	return user, nil
}

func createNewUser(
	conn *gorm.DB,
	user *UserProps,
	c echo.Context,
) (*UserProps, error) {
	if err := c.Bind(&user); err != nil {
		return &UserProps{}, err
	}

	result := conn.Debug().Create(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &UserProps{}, result.Error
	}

	return user, nil
}
