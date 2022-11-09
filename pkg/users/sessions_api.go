package users

import (
	appUtils "github.com/frankmeza/roomchat/pkg/app_utils"
)

type SessionsAPI struct {
	apiType string
}

func UseSessionsAPI() SessionsAPI {
	return SessionsAPI{apiType: "sessions"}
}

func (api SessionsAPI) CreateUserSession(user User) UserSession {
	return UserSession{
		UserSessionProps: UserSessionProps{
			Location: user.UserProps.Location,
			UserUuid: user.Uuid,
		},
		Uuid: appUtils.CreateUuid(),
	}
}
