package users

import (
	appUtils "github.com/frankmeza/roomchat/pkg/app_utils"
	"github.com/frankmeza/roomchat/pkg/errata"
	"github.com/frankmeza/roomchat/pkg/users/auth"
)

func handleSignUpMacro(user *User) error {
	uuidString := appUtils.CreateUuid()

	passwordHash, err := auth.GeneratePasswordString(
		user.UserProps.Password,
	)

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

func handleLoginMacro(user User, params handleLoginParams) (
	handleLoginMacroMetadata, error,
) {
	getUserParams := createGetUserParams(params)

	err := UseUsersAPI().GetUserByParam(&user, getUserParams)
	if err != nil {
		return handleLoginMacroMetadata{},
			errata.CreateError(err, errata.ErrMessage{
				Text: "handleLoginMacro GetUserByParam",
			})
	}

	err = auth.CheckPasswordHash(auth.CheckPasswordHashParams{
		Hash:     user.UserProps.Password,
		Password: params.Password,
	})

	if err != nil {
		return handleLoginMacroMetadata{},
			errata.CreateError(err, errata.ErrMessage{
				Text: "handleLoginMacro CheckPasswordHash doesn't match",
			})
	}

	tokenString, err := auth.GenerateTokenString(
		auth.GenerateTokenStringParams{
			Password: params.Password,
			Username: params.Username,
		})

	if err != nil {
		return handleLoginMacroMetadata{},
			errata.CreateError(err, errata.ErrMessage{
				Text: "handleLoginMacro GeneratePasswordString",
			})
	}

	var userSession UserSession
	isOk := UseSessionsAPI().CreateUserSession(user, &userSession)
	if !isOk {
		return handleLoginMacroMetadata{},
			errata.CreateError(err, errata.ErrMessage{
				Text: "handleLoginMacro GeneratePasswordString",
			})
	}

	err = UseSessionsAPI().SaveUserSession(userSession)
	if err != nil {
		return handleLoginMacroMetadata{},
			errata.CreateError(err, errata.ErrMessage{
				Text: "handleLoginMacro SaveUserSession",
			})
	}

	return handleLoginMacroMetadata{
		session: userSession,
		token:   tokenString,
	}, nil
}
