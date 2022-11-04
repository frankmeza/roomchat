package users

import (
	"github.com/frankmeza/roomchat/pkg/errata"
	jsonMap "github.com/mitchellh/mapstructure"
)

type UsersAPI struct {
	apiType string
}

func UseUsersAPI() UsersAPI {
	return UsersAPI{apiType: "users"}
}

func (api UsersAPI) CreateUser(
	user *User,
	userPropsPayload *UserProps,
	passwordHash string,
	uuid string,
) error {
	user.Uuid = uuid

	userPropsPayload.Uuid = uuid
	userPropsPayload.Password = string(passwordHash)

	err := jsonMap.Decode(userPropsPayload, &user.UserProps)
	if err != nil {
		return errata.CreateError("CreateUser jsonMap.Decode", err)
	}

	return nil
}

func (api UsersAPI) SaveUser(user *User) error {
	return saveUserDb(user)
}

func (api UsersAPI) UpdateUser(user *User) error {
	return saveUserDb(user)
}

func (api UsersAPI) GetUserByParam(user *User, params GetUserParams) error {
	return getUserDbByParam(user, params)
}
