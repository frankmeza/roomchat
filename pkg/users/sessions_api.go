package users

import (
	appUtils "github.com/frankmeza/roomchat/pkg/app_utils"
)

func UseSessionsAPI() SessionsAPI {
	return SessionsAPI{apiType: "sessions"}
}

func (api SessionsAPI) CreateUserSession(user User, session *UserSession) bool {
	*session = UserSession{
		UserSessionProps: UserSessionProps{
			Location: user.UserProps.Location,
			UserUuid: user.Uuid,
		},
		Uuid: appUtils.CreateUuid(),
	}

	return true
}

func (api SessionsAPI) SaveUserSession(session UserSession) error {
	return saveUserSessionDb(&session)
}
