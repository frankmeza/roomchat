package users

import (
	"github.com/frankmeza/roomchat/pkg/auth"
	"github.com/frankmeza/roomchat/pkg/constants"
	"github.com/frankmeza/roomchat/pkg/errata"
	"github.com/twinj/uuid"
)

func handleSignUpMacro(user *User, userProps *UserProps) error {
	uuidString := uuid.NewV4().String()

	passwordHash, err := auth.GeneratePasswordString(userProps.Password)
	if err != nil {
		return errata.CreateError("handleSignUpMacro auth.GeneratePasswordString", err)
	}

	err = UseUsersAPI().CreateUser(
		user,
		userProps,
		string(passwordHash),
		uuidString,
	)

	if err != nil {
		return errata.CreateError("handleSignUpMacro UseUsersAPI().CreateUser", err)
	}

	err = UseUsersAPI().SaveUser(user)
	if err != nil {
		return errata.CreateError("handleSignUpMacro UseUsersAPI().SaveUser", err)
	}

	return nil
}

func handleLoginMacro(user *User, params handleLoginParams) (string, error) {
	getUserParams := GetUserParams{
		Email:    params.Email,
		Username: params.Username,
	}

	if getUserParams.Email != "" {
		getUserParams.ParamName = constants.EMAIL
	}

	if getUserParams.Username != "" {
		getUserParams.ParamName = constants.USERNAME
	}

	err := UseUsersAPI().GetUserByParam(user, getUserParams)
	if err != nil {
		return "", errata.CreateError("handleLoginMacro getUserDbByParam", err)
	}

	doesPasswordMatch := auth.CheckPasswordHash(auth.CheckPasswordHashParams{
		Hash:     user.UserProps.Password,
		Password: params.Password,
	})

	if !doesPasswordMatch {
		return "", errata.CreateError("handleLoginMacro auth.CheckPasswordHash doesn't match", err)
	}

	tokenString, err := auth.GenerateTokenString(auth.GenerateTokenStringParams{
		Password: params.Password,
		Username: params.Username,
	})

	if err != nil {
		return "", errata.CreateError("handleLoginMacro auth.GeneratePasswordString", err)
	}

	return tokenString, nil
}
