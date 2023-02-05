package users

import (
	appUtils "github.com/frankmeza/roomchat/pkg/app_utils"
	"github.com/frankmeza/roomchat/pkg/auth"
	"github.com/frankmeza/roomchat/pkg/errata"
	"github.com/frankmeza/roomchat/pkg/sessions"
)

func handleSignUpMacro(user *User) error {
	uuidString := appUtils.CreateUuid()

	passwordHash, err := auth.GeneratePasswordString(
		user.UserProps.Password,
	)

	if err != nil {
		return errata.CreateError(err, []string{
			"handleSignUpMacro GeneratePasswordString",
		})
	}

	err = UseUsersAPI().CreateUser(user, CreateUserParams{
		Hash: string(passwordHash),
		Uuid: uuidString,
	})

	if err != nil {
		return errata.CreateError(err, []string{
			"handleSignUpMacro CreateUser",
		})
	}

	err = UseUsersAPI().SaveUser(user)
	if err != nil {
		return errata.CreateError(err, []string{
			"handleSignUpMacro SaveUser",
		})
	}

	return nil
}

func handleLoginMacro(
	user User,
	params loginParams,
) (loginMacroMetadata, error) {
	getUserParams := createGetUserParams(params)

	err := UseUsersAPI().GetUserByParam(&user, getUserParams)
	if err != nil {
		return loginMacroMetadata{},
			errata.CreateError(err, []string{
				"handleLoginMacro GetUserByParam",
			})
	}

	err = auth.CheckPasswordHash(auth.CheckPasswordHashParams{
		Hash:     user.UserProps.Password,
		Password: params.Password,
	})

	if err != nil {
		return loginMacroMetadata{},
			errata.CreateError(err, []string{
				"handleLoginMacro CheckPasswordHash doesn't match",
			})
	}

	tokenString, err := auth.GenerateTokenString(
		auth.GenerateTokenStringParams{
			Password: params.Password,
			Username: params.Username,
		})

	if err != nil {
		return loginMacroMetadata{},
			errata.CreateError(err, []string{
				"handleLoginMacro GeneratePasswordString",
			})
	}

	var session sessions.Session

	sessionParams := sessions.SessionParams{
		Location: user.UserProps.Location,
		Uuid:     user.UserProps.Uuid,
	}

	isOk := sessions.UseSessionsAPI().CreateSession(
		sessionParams,
		&session,
	)

	if !isOk {
		return loginMacroMetadata{},
			errata.CreateError(err, []string{
				"handleLoginMacro GeneratePasswordString",
			})
	}

	err = sessions.UseSessionsAPI().SaveSession(session)
	if err != nil {
		return loginMacroMetadata{},
			errata.CreateError(err, []string{
				"handleLoginMacro SaveSession",
			})
	}

	return loginMacroMetadata{
		session: session,
		token:   tokenString,
	}, nil
}
