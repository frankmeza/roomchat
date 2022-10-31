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
		return errata.CreateError(errata.ErrataParams{
			Err:    err,
			ErrMsg: "handleSignUpMacro auth.GeneratePasswordString",
		})
	}

	err = UseUsersAPI().CreateUser(
		user,
		userProps,
		string(passwordHash),
		uuidString,
	)

	if err != nil {
		return errata.CreateError(errata.ErrataParams{
			Err:    err,
			ErrMsg: "handleSignUpMacro UseUsersAPI().CreateUser",
		})
	}

	err = UseUsersAPI().SaveUser(user)
	if err != nil {
		return errata.CreateError(errata.ErrataParams{
			Err:    err,
			ErrMsg: "handleSignUpMacro UseUsersAPI().SaveUser",
		})
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
		return "", errata.CreateError(errata.ErrataParams{
			Err:    err,
			ErrMsg: "handleLoginMacro getUserDbByParam",
		})
	}

	doesPasswordMatch := auth.CheckPasswordHash(auth.CheckPasswordHashParams{
		Hash:     user.UserProps.Password,
		Password: params.Password,
	})

	if !doesPasswordMatch {
		return "", errata.CreateError(errata.ErrataParams{
			Err:    err,
			ErrMsg: "handleLoginMacro auth.CheckPasswordHash doesn't match",
		})
	}

	tokenString, err := auth.GenerateTokenString(auth.GenerateTokenStringParams{
		Password: params.Password,
		Username: params.Username,
	})

	if err != nil {
		return "", errata.CreateError(errata.ErrataParams{
			Err:    err,
			ErrMsg: "handleLoginMacro auth.GeneratePasswordString",
		})
	}

	return tokenString, nil
}
