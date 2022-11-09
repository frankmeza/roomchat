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
		return errata.CreateError(err, errata.ErrMessage{
			Text: "saveUserDb GetDbConnection",
		})
	}

	result := dbConn.Debug().Create(user)
	if result.Error != nil {
		return errata.CreateError(result.Error, errata.ErrMessage{
			Text: "saveUserDb Create",
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
		return params.Uuid
	}

	return ""
}

func getUserDbByParam(user *User, params GetUserParams) error {
	dbConn, err := db.GetDbConnection()
	if err != nil {
		return errata.CreateError(err, errata.ErrMessage{
			Text: "getUserDbByParam GetDbConnection",
		})
	}

	paramToUse := getParamToUse(params)
	if paramToUse == "" {
		return errata.CreateError(err, errata.ErrMessage{
			Text: "getUserDbByParam getParamToUse",
		})
	}

	query := datatypes.
		JSONQuery(constants.USER_PROPS).
		Equals(paramToUse, params.ParamName)

	result := dbConn.Debug().Find(user, query)
	if result.Error != nil {
		return errata.CreateError(result.Error, errata.ErrMessage{
			Text: "getUserDbByParam Find",
		})
	}

	return nil
}
