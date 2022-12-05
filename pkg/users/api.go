package users

import (
	"github.com/frankmeza/roomchat/pkg/errata"
	jsonMap "github.com/mitchellh/mapstructure"
)

func UseUsersAPI() UsersAPI {
	return UsersAPI{apiType: "users"}
}

type CreateUserParams struct {
	Hash string
	Uuid string
}

func (api UsersAPI) CreateUser(user *User, params CreateUserParams) error {
	user.Uuid = params.Uuid

	user.UserProps.Uuid = params.Uuid
	user.UserProps.Password = string(params.Hash)

	err := jsonMap.Decode(user.UserProps, &user.UserProps)
	if err != nil {
		return errata.CreateError(err, []string{
			"CreateUser Decode",
		})
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
