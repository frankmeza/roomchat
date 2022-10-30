package users

import (
	"net/http"

	"github.com/frankmeza/roomchat/pkg/auth"
	cc "github.com/frankmeza/roomchat/pkg/constants"
	"github.com/frankmeza/roomchat/pkg/errata"
	"github.com/frankmeza/roomchat/pkg/utils"
	"github.com/labstack/echo/v4"
)

func handleGetUsers(context echo.Context) error {
	var users []User
	actionGetUsers(&users)

	return context.JSON(http.StatusOK, users)
}

func handleGetUserByUsername(context echo.Context) error {
	username := context.Param(cc.USERNAME)

	foundUser, err := actionGetUserByUsername(username, cc.NO_PASSWORD, false)
	if err != nil {
		return errata.HandlerError(errata.HandlerErrorParams{
			CallingFn: "actionGetUserByUsername",
			Context:   context,
			Err:       err,
			Status:    http.StatusBadRequest,
		})
	}

	return context.JSON(http.StatusOK, foundUser)
}

func handleSignUp(context echo.Context) error {
	var userPropsPayload UserProps

	err := context.Bind(&userPropsPayload)
	if err != nil {
		errorString := "handleSignUp error: " + err.Error()
		return context.String(http.StatusBadRequest, errorString)
	}

	savedUser, err := actionCreateUser(&userPropsPayload)
	if err != nil {
		errorString := "actionCreateUser error: " + err.Error()
		return context.String(http.StatusBadRequest, errorString)
	}

	return context.JSON(http.StatusCreated, savedUser)
}

// todo - add support for email along with username
func handleLogin(context echo.Context) error {
	username := context.Param(cc.USERNAME)
	password := context.Param(cc.PASSWORD)

	foundUser, err := actionGetUserByUsername(username, password, true)
	if err != nil {
		return utils.ReturnError("actionGetUserByUsername", err)
	}

	doesPasswordMatch := auth.CheckPasswordHash(
		foundUser.UserProps.Password,
		password,
	)

	if !doesPasswordMatch {
		return context.String(
			http.StatusBadRequest,
			cc.LOGIN_ERROR,
		)
	}

	tokenString, err := auth.GeneratePasswordString(password)
	if err != nil {
		return context.String(
			http.StatusBadRequest,
			cc.LOGIN_ERROR,
		)
	}

	return context.JSON(http.StatusOK, echo.Map{
		"token": tokenString,
	})
}
