package users

import (
	"github.com/frankmeza/roomchat/pkg/db"
	"github.com/frankmeza/roomchat/pkg/errata"
)

func saveUserSessionDb(session *UserSession) error {
	dbConn, err := db.GetDbConnection()
	if err != nil {
		return errata.CreateError(err, []string{
			"saveUserSessionDb GetDbConnection",
		})
	}

	result := dbConn.Debug().Create(session)
	if result.Error != nil {
		return errata.CreateError(result.Error, []string{
			"saveUserSessionDb Create",
		})
	}

	return nil
}
