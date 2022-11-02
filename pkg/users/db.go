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
		return errata.CreateError("saveUserDb db.GetDbConnection", err)
	}

	result := dbConn.Debug().Create(user)
	if result.Error != nil {
		return errata.CreateError("saveUserDb dbConn.Debug().Create", err)
	}

	return nil
}

type GetUserParams struct {
	Email     string
	ID        string
	ParamName string
	Username  string
	Uuid      string
}

func getParamToUse(params GetUserParams) string {
	if params.ParamName == constants.EMAIL {
		return params.Email
	}

	if params.ParamName == constants.ID {
		return params.ID
	}

	if params.ParamName == constants.USERNAME {
		return params.Username
	}

	if params.ParamName == constants.UUID {
		return params.Uuid
	}

	return ""
}

func getUserDbByParam(user *User, params GetUserParams) error {
	dbConn, err := db.GetDbConnection()
	if err != nil {
		return errata.CreateError("getUserDbByParam db.GetDbConnection", err)
	}

	paramToUse := getParamToUse(params)
	if paramToUse == "" {
		return errata.CreateError("getUserDbByParam getParamToUse", err)
	}

	query := datatypes.
		JSONQuery(constants.USER_PROPS).
		Equals(paramToUse, params.ParamName)

	result := dbConn.Debug().Find(&user, query)
	if result.Error != nil {
		return errata.CreateError("getUserDbByParam db.GetDbConnection", err)
	}

	return nil
}
