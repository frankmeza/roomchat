package connections

import (
	"github.com/frankmeza/roomchat/pkg/db"
	"github.com/frankmeza/roomchat/pkg/errata"
)

func saveConnectionDb(connection *Connection) error {
	dbConn, err := db.GetDbConnection()
	if err != nil {
		return errata.CreateError(errata.ErrataParams{
			Err:     err,
			ErrFunc: "saveConnectionDb db.GetDbConnection",
		})
	}

	result := dbConn.Debug().Create(connection)
	if result.Error != nil {
		return errata.CreateError(errata.ErrataParams{
			Err:     result.Error,
			ErrFunc: "saveConnectionDb dbConn.Debug().Create",
		})
	}

	return nil
}
