package connections

import (
	"errors"

	"github.com/frankmeza/roomchat/pkg/constants"
	"github.com/frankmeza/roomchat/pkg/errata"
	"github.com/frankmeza/roomchat/pkg/users"
)

func handleMakeConnectionMacro(connection *Connection) error {
	var verifiedUsers []users.User

	userUuids := []string{connection.FromUser, connection.ToUser}
	for _, userUuid := range userUuids {
		var verifiedUser users.User

		users.UseUsersAPI().GetUserByParam(&verifiedUser, users.GetUserParams{
			ParamName: constants.UUID,
			Uuid:      userUuid,
		})

		verifiedUsers = append(verifiedUsers, verifiedUser)
	}

	if len(verifiedUsers) != 2 {
		return errata.CreateError(errata.ErrataParams{
			Err:    errors.New("length of verified users is wrong"),
			ErrMsg: "handleMakeConnectionMacro",
		})
	}

	err := useConnectionsAPI().SaveConnection(connection)
	if err != nil {
		return errata.CreateError(errata.ErrataParams{
			Err:    err,
			ErrMsg: "handleMakeConnectionMacro SaveConnection",
		})
	}

	return nil
}
