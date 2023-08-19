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
		return errata.CreateError(err, []string{
			"saveUserDb GetDbConnection",
		})
	}

	result := dbConn.Debug().Create(user)
	if result.Error != nil {
		return errata.CreateError(result.Error, []string{
			"saveUserDb Create",
		})
	}

	return nil
}

func getUserDbByParam(user *User, params GetUserParams) (UserProps, error) {
	dbConn, err := db.GetDbConnection()
	if err != nil {
		return UserProps{}, errata.CreateError(err, []string{
			"getUserDbByParam GetDbConnection",
		})
	}

	paramToUse := getParamToUse(params)
	if paramToUse == "" {
		return UserProps{}, errata.CreateError(err, []string{
			"getUserDbByParam getParamToUse",
		})
	}

	query := datatypes.
		JSONQuery(constants.USER_PROPS).
		Equals(paramToUse, params.ParamName)

	result := dbConn.Debug().Find(user, query)
	if result.Error != nil {
		return UserProps{}, errata.CreateError(result.Error, []string{
			"getUserDbByParam Find",
		})
	}

	return user.UserProps, nil
}
