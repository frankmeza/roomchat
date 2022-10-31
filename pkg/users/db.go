package users

import (
	"github.com/frankmeza/roomchat/pkg/constants"
	"github.com/frankmeza/roomchat/pkg/db"
	"github.com/frankmeza/roomchat/pkg/errata"
	"gorm.io/datatypes"
)

func saveUserDb(user *User) error {
	dbConn, err := db.GetDbConnection()
	if err != nil {
		return errata.CreateError(errata.ErrataParams{
			Err:     err,
			ErrFunc: "saveUserDb db.GetDbConnection",
		})
	}

	result := dbConn.Debug().Create(user)
	if result.Error != nil {
		return errata.CreateError(errata.ErrataParams{
			Err:     result.Error,
			ErrFunc: "saveUserDb dbConn.Debug().Create",
		})
	}

	return nil
}

type GetUserParams struct {
	Email     string
	ID        string
	ParamName string
	Username  string
}

func getUserDbByParam(user *User, params GetUserParams) error {
	dbConn, err := db.GetDbConnection()
	if err != nil {
		return errata.CreateError(errata.ErrataParams{
			Err:     err,
			ErrFunc: "getUserDbByParam db.GetDbConnection",
		})
	}

	var paramToUse interface{}

	if params.ParamName == constants.EMAIL {
		paramToUse = params.Email
	}

	if params.ParamName == constants.ID {
		paramToUse = params.ID
	}

	if params.ParamName == constants.USERNAME {
		paramToUse = params.Username
	}

	query := datatypes.
		JSONQuery(constants.USER_PROPS).
		Equals(paramToUse, params.ParamName)

	result := dbConn.Debug().Find(&user, query)

	if result.Error != nil {
		return errata.CreateError(errata.ErrataParams{
			Err:     result.Error,
			ErrFunc: "getUserDbByParam db.GetDbConnection",
		})
	}

	return nil
}
