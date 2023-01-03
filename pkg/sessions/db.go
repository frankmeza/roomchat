package sessions

import (
	"github.com/frankmeza/roomchat/pkg/db"
	"github.com/frankmeza/roomchat/pkg/errata"
)

func saveSessionDb(session *Session) error {
	dbConn, err := db.GetDbConnection()
	if err != nil {
		return errata.CreateError(err, []string{
			"saveSessionDb GetDbConnection",
		})
	}

	result := dbConn.Debug().Create(session)
	if result.Error != nil {
		return errata.CreateError(result.Error, []string{
			"saveSessionDb Create",
		})
	}

	return nil
}
