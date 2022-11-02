package connections

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

func (api ConnectionsAPI) SaveConnection(connection *Connection) error {
	return saveConnectionDb(connection)
}

func (api ConnectionsAPI) GetConnectionByParam(
	connection *Connection,
	params getConnectionParams,
) error {
	return getConnectionDbByParam(connection, params)
}
