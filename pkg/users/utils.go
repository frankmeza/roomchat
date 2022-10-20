package users

import (
	"errors"

	cc "github.com/frankmeza/roomchat/pkg/constants"
	"github.com/frankmeza/roomchat/pkg/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type GetUserParams struct {
	Email     string
	ID        string
	ParamName string
}

func getUserById(
	conn *gorm.DB,
	user *User,
	id string,
) error {
	result := conn.First(&user, "id = ?", id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}

	return result.Error
}

func getUserByParam(
	conn *gorm.DB,
	user *User,
	params GetUserParams,
) error {
	var paramToUse interface{}
	if params.ParamName == cc.EMAIL {
		paramToUse = params.Email
	}

	if params.ParamName == cc.ID {
		getUserById(conn, user, params.ID)
	}

	query := datatypes.
		JSONQuery(cc.USER_SPECS).
		Equals(paramToUse, params.ParamName)

	result := conn.Debug().Find(&user, query)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}

	return nil
}

func createNewUser(
	conn *gorm.DB,
	user *User,
	c echo.Context,
) error {
	if err := c.Bind(&user); err != nil {
		return err
	}

	result := conn.Debug().Create(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}

	return nil
}

func deleteUser(
	conn *gorm.DB,
	user *User,
	c echo.Context,
) error {
	result := conn.Debug().Delete(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return utils.ReturnError("conn.Debug().Delete", result.Error)
	}

	return nil
}
