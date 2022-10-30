package users

import (
	"github.com/frankmeza/roomchat/pkg/auth"
	"github.com/frankmeza/roomchat/pkg/constants"
	"github.com/frankmeza/roomchat/pkg/errata"
	"github.com/twinj/uuid"
)

func handleSignUpMacro(user *User, userPropsPayload *UserProps) error {
	uuidString := uuid.NewV4().String()

	passwordHash, err := auth.GeneratePasswordString(userPropsPayload.Password)
	if err != nil {
		return errata.CreateError(errata.ErrataParams{
			Err:     err,
			ErrFunc: "handleSignUpMacro auth.GeneratePasswordString",
		})
	}

	err = useUsersAPI().CreateUser(
		user,
		userPropsPayload,
		string(passwordHash),
		uuidString,
	)

	if err != nil {
		return errata.CreateError(errata.ErrataParams{
			Err:     err,
			ErrFunc: "handleSignUpMacro useUsersAPI().CreateUser",
		})
	}

	err = useUsersAPI().SaveUser(user)
	if err != nil {
		return errata.CreateError(errata.ErrataParams{
			Err:     err,
			ErrFunc: "handleSignUpMacro useUsersAPI().SaveUser",
		})
	}

	return nil
}

func handleLoginMacro(user *User, username, password string) (string, error) {
	params := GetUserParams{Username: username, ParamName: constants.USERNAME}

	err := useUsersAPI().GetUserByParam(user, params)
	if err != nil {
		return "", errata.CreateError(errata.ErrataParams{
			Err:     err,
			ErrFunc: "handleLoginMacro getUserDbByParam",
		})
	}

	doesPasswordMatch := auth.CheckPasswordHash(auth.CheckPasswordHashParams{
		Hash:     user.UserProps.Password,
		Password: password,
	})

	if !doesPasswordMatch {
		return "", errata.CreateError(errata.ErrataParams{
			Err:     err,
			ErrFunc: "handleLoginMacro auth.CheckPasswordHash doesn't match",
		})
	}

	tokenString, err := auth.GenerateTokenString(auth.GenerateTokenStringParams{
		Password: password,
		Username: username,
	})

	if err != nil {
		return "", errata.CreateError(errata.ErrataParams{
			Err:     err,
			ErrFunc: "handleLoginMacro auth.GeneratePasswordString",
		})
	}

	return tokenString, nil
}
