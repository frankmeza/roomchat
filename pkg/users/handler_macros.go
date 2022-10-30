package users

import (
	"github.com/frankmeza/roomchat/pkg/auth"
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
