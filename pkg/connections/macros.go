package connections

import (
	"errors"

	"github.com/frankmeza/roomchat/pkg/constants"
	"github.com/frankmeza/roomchat/pkg/errata"
	"github.com/frankmeza/roomchat/pkg/users"
)

func handleMakeConnectionMacro(
	params createConnectionParams,
) (Connection, error) {
	var verifiedUsers []users.User

	userUuids := []string{
		params.Message.FromUser,
		params.Message.ToUser,
	}

	for _, userUuid := range userUuids {
		var verifiedUser users.User

		users.UseUsersAPI().GetUserByParam(
			&verifiedUser, users.GetUserParams{
				ParamName: constants.UUID,
				Uuid:      userUuid,
			},
		)

		verifiedUsers = append(verifiedUsers, verifiedUser)
	}

	if len(verifiedUsers) != 2 {
		return Connection{}, errata.CreateError(errors.New(""), []string{
			"handleMakeConnectionMacro",
		})
	}

	err := useConnectionsAPI().SaveMessage(&params.Message)
	if err != nil {
		return Connection{}, errata.CreateError(err, []string{
			"handleMakeConnectionMacro SaveMessage",
		})
	}

	connection := useConnectionsAPI().CreateConnection(params)
	err = useConnectionsAPI().SaveConnection(&connection)
	if err != nil {
		return Connection{}, errata.CreateError(err, []string{
			"handleMakeConnectionMacro SaveConnection",
		})
	}

	return connection, nil
}

func handleAddMessageMacro(params addMessageParams) error {
	var connection Connection

	err := useConnectionsAPI().GetConnectionByParam(
		&connection, getConnectionParams{
			ParamName: constants.UUID,
			Uuid:      params.ConnectionUuid,
		},
	)

	if err != nil {
		return errata.CreateError(err, []string{
			"handleAddMessageMacro GetConnectionByParam",
		})
	}

	var toUser string

	if params.FromUser == connection.ConnectionProps.FromUser {
		toUser = connection.ConnectionProps.ToUser
	} else {
		toUser = connection.ConnectionProps.FromUser
	}

	newMessage := Message{FromUser: params.FromUser, ToUser: toUser}

	err = useConnectionsAPI().SaveMessage(&newMessage)
	if err != nil {
		return errata.CreateError(err, []string{
			"handleAddMessageMacro SaveMessage",
		})
	}

	connection.Messages = append(connection.Messages, newMessage)

	err = useConnectionsAPI().UpdateConnection(&connection)
	if err != nil {
		return errata.CreateError(err, []string{
			"handleAddMessageMacro UpdateConnection",
		})
	}

	return nil
}
