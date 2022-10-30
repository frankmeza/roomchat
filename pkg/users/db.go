package users

import (
	"github.com/frankmeza/roomchat/pkg/db"
	"github.com/frankmeza/roomchat/pkg/errata"
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

// import (
// 	"errors"

// 	cc "github.com/frankmeza/roomchat/pkg/constants"
// 	"github.com/frankmeza/roomchat/pkg/db"
// 	"gorm.io/datatypes"
// 	"gorm.io/gorm"
// )

// type GetUserParams struct {
// 	Email     string
// 	ID        string
// 	ParamName string
// 	Username  string
// }

// func getUsersDb(dbConn *gorm.DB, users *[]User) error {
// 	result := dbConn.First(&users)
// 	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 		return result.Error
// 	}

// 	return nil
// }

// func getUserById(dbConn *gorm.DB, user *User, id string) error {
// 	result := dbConn.First(&user, id)
// 	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 		return result.Error
// 	}

// 	return result.Error
// }

// func getUserByParam(dbConn *gorm.DB, user *User, params GetUserParams) error {
// 	var paramToUse interface{}
// 	if params.ParamName == cc.EMAIL {
// 		paramToUse = params.Email
// 	}

// 	if params.ParamName == cc.ID {
// 		getUserById(dbConn, user, params.ID)
// 	}

// 	if params.ParamName == cc.USERNAME {
// 		paramToUse = params.Username
// 	}

// 	query := datatypes.
// 		JSONQuery(cc.USER_PROPS).
// 		Equals(paramToUse, params.ParamName)

// 	result := dbConn.Debug().Find(&user, query)
// 	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 		return result.Error
// 	}

// 	return nil
// }
