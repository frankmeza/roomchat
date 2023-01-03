package sessions

import (
	appUtils "github.com/frankmeza/roomchat/pkg/app_utils"
)

func UseSessionsAPI() SessionsAPI {
	return SessionsAPI{apiType: "sessions"}
}

func (api SessionsAPI) CreateSession(
	params CreateSessionParams,
	session *Session,
) bool {
	*session = Session{
		SessionProps: SessionProps{
			Location: params.Location,
			UserUuid: params.Uuid,
		},
		Uuid: appUtils.CreateUuid(),
	}

	return true
}

func (api SessionsAPI) SaveSession(session Session) error {
	return saveSessionDb(&session)
}
