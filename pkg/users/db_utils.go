package users

import (
	"errors"

	cc "github.com/frankmeza/roomchat/pkg/constants"
	"github.com/frankmeza/roomchat/pkg/db"
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
	dbConn *gorm.DB,
	user *User,
	id string,
) error {
	result := dbConn.First(&user, "select * where id = ?", id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}

	return result.Error
}

func getUserByParam(
	dbConn *gorm.DB,
	user *User,
	params GetUserParams,
) error {
	var paramToUse interface{}
	if params.ParamName == cc.EMAIL {
		paramToUse = params.Email
	}

	if params.ParamName == cc.ID {
		getUserById(dbConn, user, params.ID)
	}

	query := datatypes.
		JSONQuery(cc.USER_SPECS).
		Equals(paramToUse, params.ParamName)

	result := dbConn.Debug().Find(&user, query)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}

	return nil
}

func saveUserDb(c echo.Context, user *User) error {
	dbConn := db.GetDbConnection()

	result := dbConn.Debug().Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
