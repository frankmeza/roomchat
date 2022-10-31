package connections

type ConnectionsAPI struct {
	apiType string
}

func useConnectionsAPI() ConnectionsAPI {
	return ConnectionsAPI{apiType: "connections"}
}

func (api ConnectionsAPI) SaveConnection(connection *Connection) error {
	return saveConnectionDb(connection)
}
