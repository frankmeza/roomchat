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
			Err:    err,
			ErrMsg: "saveUserDb db.GetDbConnection",
		})
	}

	result := dbConn.Debug().Create(user)
	if result.Error != nil {
		return errata.CreateError(errata.ErrataParams{
			Err:    result.Error,
			ErrMsg: "saveUserDb dbConn.Debug().Create",
		})
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
		return params.Username
	}

	return ""
}

func getUserDbByParam(user *User, params GetUserParams) error {
	dbConn, err := db.GetDbConnection()
	if err != nil {
		return errata.CreateError(errata.ErrataParams{
			Err:    err,
			ErrMsg: "getUserDbByParam db.GetDbConnection",
		})
	}

	paramToUse := getParamToUse(params)
	if paramToUse == "" {
		return errata.CreateError(errata.ErrataParams{
			Err:    err,
			ErrMsg: "getUserDbByParam getParamToUse",
		})
	}

	query := datatypes.
		JSONQuery(constants.USER_PROPS).
		Equals(paramToUse, params.ParamName)

	result := dbConn.Debug().Find(&user, query)
	if result.Error != nil {
		return errata.CreateError(errata.ErrataParams{
			Err:    result.Error,
			ErrMsg: "getUserDbByParam db.GetDbConnection",
		})
	}

	return nil
}
