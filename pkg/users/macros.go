package users

import (
	"github.com/frankmeza/roomchat/pkg/errata"
	"github.com/frankmeza/roomchat/pkg/users/auth"
	"github.com/twinj/uuid"
)

func handleSignUpMacro(user *User) error {
	uuidString := uuid.NewV4().String()

	passwordHash, err := auth.GeneratePasswordString(user.UserProps.Password)
	if err != nil {
		return errata.CreateError(err, errata.ErrMessage{
			Text: "handleSignUpMacro GeneratePasswordString",
		})
	}

	err = UseUsersAPI().CreateUser(user, CreateUserParams{
		Hash: string(passwordHash),
		Uuid: uuidString,
	})

	if err != nil {
		return errata.CreateError(err, errata.ErrMessage{
			Text: "handleSignUpMacro CreateUser",
		})
	}

	err = UseUsersAPI().SaveUser(user)
	if err != nil {
		return errata.CreateError(err, errata.ErrMessage{
			Text: "handleSignUpMacro SaveUser",
		})
	}

	return nil
}

func handleLoginMacro(user *User, params handleLoginParams) (string, error) {
	getUserParams := createGetUserParams(params)

	err := UseUsersAPI().GetUserByParam(user, getUserParams)
	if err != nil {
		return "", errata.CreateError(err, errata.ErrMessage{
			Text: "handleLoginMacro GetUserByParam",
		})
	}

	err = auth.CheckPasswordHash(auth.CheckPasswordHashParams{
		Hash:     user.UserProps.Password,
		Password: params.Password,
	})

	if err != nil {
		return "", errata.CreateError(err, errata.ErrMessage{
			Text: "handleLoginMacro auth.CheckPasswordHash doesn't match",
		})
	}

	tokenString, err := auth.GenerateTokenString(auth.GenerateTokenStringParams{
		Password: params.Password,
		Username: params.Username,
	})

	if err != nil {
		return "", errata.CreateError(err, errata.ErrMessage{
			Text: "handleLoginMacro auth.GeneratePasswordString",
		})
	}

	userSession := UseSessionsAPI().CreateUserSession(*user)
	// now save user session,
	// return with tokenString?

	return tokenString, nil
}
