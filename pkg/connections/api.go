package connections

import (
	appUtils "github.com/frankmeza/roomchat/pkg/app_utils"
)

type ConnectionsAPI struct {
	apiType string
}

func useConnectionsAPI() ConnectionsAPI {
	return ConnectionsAPI{apiType: "connections"}
}

func (api ConnectionsAPI) SaveMessage(message *Message) error {
	return saveMessageDb(message)
}

func (api ConnectionsAPI) UpdateConnection(connection *Connection) error {
	return updateConnectionDb(connection)
}

func (api ConnectionsAPI) CreateConnection(
	params handleMakeConnectionParams,
) Connection {
	uuidString := appUtils.CreateUuid()

	return Connection{
		ConnectionProps: ConnectionProps{
			FromUser: params.Message.FromUser,
			Location: params.Location,
			ToUser:   params.Message.ToUser,
			Uuid:     uuidString,
		},
		Messages: []Message{params.Message},
		Uuid:     uuidString,
	}
}

func (api ConnectionsAPI) SaveConnection(connection *Connection) error {
	return saveConnectionDb(connection)
}

func (api ConnectionsAPI) GetConnectionByParam(
	connection *Connection,
	params getConnectionParams,
) error {
	return getConnectionDbByParam(connection, params)
}
