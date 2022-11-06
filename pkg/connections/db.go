package connections

import (
	"fmt"

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
		return errata.CreateError("updateConnectionDb GetConnectionByParam", err)
	}

	if existingConnection.Uuid != updatedConnection.Uuid {
		return errata.CreateError("cannot update connection with this id", err)
	}

	// save new state

	return nil
}

func saveConnectionDb(connection *Connection) error {
	dbConn, err := db.GetDbConnection()
	if err != nil {
		return errata.CreateError("saveConnectionDb GetDbConnection", err)
	}

	result := dbConn.Debug().Create(connection)
	if result.Error != nil {
		return errata.CreateError("saveConnectionDb Create", result.Error)
	}

	return nil
}

func saveMessageDb(Message *Message) error {
	dbConn, err := db.GetDbConnection()
	if err != nil {
		return errata.CreateError("saveMessageDb GetDbMessage", err)
	}

	result := dbConn.Debug().Create(Message)
	if result.Error != nil {
		return errata.CreateError("saveMessageDb Create", result.Error)
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
		return errata.CreateError("getConnectionDbByParam GetDbConnection", err)
	}

	paramToUse := getParamToUse(params)
	if paramToUse == "" {
		return errata.CreateError("getConnectionDbByParam getParamToUse", err)
	}

	query := datatypes.
		JSONQuery(constants.CONNECTIONS).
		Equals(paramToUse, params.ParamName)

	fmt.Println("query is", query)

	result := dbConn.Debug().Find(&connection, query)
	if result.Error != nil {
		return errata.CreateError("getConnectionDbByParam Find", err)
	}

	return nil

}
