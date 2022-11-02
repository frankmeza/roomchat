package connections

import (
	"errors"

	"github.com/frankmeza/roomchat/pkg/constants"
	"github.com/frankmeza/roomchat/pkg/errata"
	"github.com/frankmeza/roomchat/pkg/users"
)

func handleMakeConnectionMacro(params handleMakeConnectionParams) (Connection, error) {
	var verifiedUsers []users.User

	userUuids := []string{params.Message.FromUser, params.Message.ToUser}
	for _, userUuid := range userUuids {
		var verifiedUser users.User

		users.UseUsersAPI().GetUserByParam(&verifiedUser, users.GetUserParams{
			ParamName: constants.UUID,
			Uuid:      userUuid,
		})

		verifiedUsers = append(verifiedUsers, verifiedUser)
	}

	if len(verifiedUsers) != 2 {
		return Connection{},
			errata.CreateError("handleMakeConnectionMacro", errors.New(""))
	}

	err := useConnectionsAPI().SaveMessage(&params.Message)
	if err != nil {
		return Connection{},
			errata.CreateError("handleMakeConnectionMacro SaveConnection", err)
	}

	connection := Connection{
		FromUser: params.Message.FromUser,
		Location: params.Location,
		Messages: []Message{params.Message},
		ToUser:   params.Message.ToUser,
	}

	err = useConnectionsAPI().SaveConnection(&connection)
	if err != nil {
		return Connection{},
			errata.CreateError("handleMakeConnectionMacro SaveConnection", err)
	}

	return connection, nil
}

func handleAddMessageMacro(params handleAddMessageParams) error {
	var connection Connection

	err := useConnectionsAPI().GetConnectionByParam(
		&connection, getConnectionParams{
			ParamName: constants.UUID,
			Uuid:      params.ConnectionUuid,
		},
	)

	if err != nil {
		return errata.CreateError("handleAddMessageMacro GetConnectionByParam", err)
	}

	toUser := connection.FromUser
	if params.FromUser == connection.FromUser {
		toUser = connection.ToUser
	}

	newMessage := Message{FromUser: params.FromUser, ToUser: toUser}
	err = useConnectionsAPI().SaveMessage(&newMessage)
	if err != nil {
		return errata.CreateError("handleAddMessageMacro SaveMessage", err)
	}

	connection.Messages = append(connection.Messages, newMessage)

	err = useConnectionsAPI().UpdateConnection(&connection)
	if err != nil {
		return errata.CreateError("handleAddMessageMacro UpdateConnection", err)
	}

	return nil
}
