package connections

import "github.com/frankmeza/roomchat/pkg/errata"

func handleMakeConnectionMacro(connection *Connection) error {
	err := useConnectionsAPI().SaveConnection(connection)
	if err != nil {
		return errata.CreateError(errata.ErrataParams{
			Err:     err,
			ErrFunc: "handleMakeConnectionMacro SaveConnection",
		})
	}

	return nil
}
