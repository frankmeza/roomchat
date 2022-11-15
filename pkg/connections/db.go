package connections

import (
	"github.com/frankmeza/roomchat/pkg/constants"
	"github.com/frankmeza/roomchat/pkg/db"
	"github.com/frankmeza/roomchat/pkg/errata"
	"gorm.io/datatypes"
)

func updateConnectionDb(updatedConnection *Connection) error {
	var existingConnection Connection

	err := useConnectionsAPI().GetConnectionByParam(
		&existingConnection,
		getConnectionParams{
			ParamName: constants.UUID,
			Uuid:      existingConnection.Uuid,
		})

	if err != nil {
		return errata.CreateError(err, errata.ErrMessage{
			Text: "updateConnectionDb GetConnectionByParam",
		})
	}

	if existingConnection.Uuid != updatedConnection.Uuid {
		return errata.CreateError(err, errata.ErrMessage{
			Text: "cannot update connection with this id",
		})
	}

	// save new state

	return nil
}

func saveConnectionDb(connection *Connection) error {
	dbConn, err := db.GetDbConnection()
	if err != nil {
		return errata.CreateError(err, errata.ErrMessage{
			Text: "saveConnectionDb GetDbConnection",
		})
	}

	result := dbConn.Debug().Create(connection)
	if result.Error != nil {
		return errata.CreateError(result.Error, errata.ErrMessage{
			Text: "saveConnectionDb Create",
		})
	}

	return nil
}

func saveMessageDb(Message *Message) error {
	dbConn, err := db.GetDbConnection()
	if err != nil {
		return errata.CreateError(err, errata.ErrMessage{
			Text: "saveMessageDb GetDbMessage",
		})
	}

	result := dbConn.Debug().Create(Message)
	if result.Error != nil {
		return errata.CreateError(result.Error, errata.ErrMessage{
			Text: "saveMessageDb Create",
		})
	}

	return nil
}

type getConnectionParams struct {
	ParamName string
	Uuid      string
}

func getParamToUse(params getConnectionParams) string {
	if params.ParamName == constants.UUID {
		return params.Uuid
	}

	return ""
}

func getConnectionDbByParam(
	connection *Connection,
	params getConnectionParams,
) error {
	dbConn, err := db.GetDbConnection()
	if err != nil {
		return errata.CreateError(err, errata.ErrMessage{
			Text: "getConnectionDbByParam GetDbConnection",
		})
	}

	paramToUse := getParamToUse(params)
	if paramToUse == "" {
		return errata.CreateError(err, errata.ErrMessage{
			Text: "getConnectionDbByParam getParamToUse",
		})
	}

	query := datatypes.
		JSONQuery(constants.CONNECTIONS).
		Equals(paramToUse, params.ParamName)

	result := dbConn.Debug().Find(&connection, query)
	if result.Error != nil {
		return errata.CreateError(result.Error, errata.ErrMessage{
			Text: "getConnectionDbByParam Find",
		})
	}

	return nil

}
