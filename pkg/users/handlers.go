package users

import (
	"net/http"

	"github.com/frankmeza/roomchat/pkg/response"
	"github.com/labstack/echo/v4"
)

func handleSignUp(context echo.Context) error {
	var user User
	var userPropsPayload UserProps

	err := context.Bind(&userPropsPayload)
	if err != nil {
		return response.HandlerError(response.HandlerErrorParams{
			Context: context,
			Err:     err,
			ErrFunc: "handleSignUp context.Bind",
			Status:  http.StatusBadRequest,
		})
	}

	err = handleSignUpMacro(&user, &userPropsPayload)
	if err != nil {
		return response.HandlerError(response.HandlerErrorParams{
			Context: context,
			Err:     err,
			ErrFunc: "handleSignUp handleSignUpMacro",
			Status:  http.StatusBadRequest,
		})
	}

	return response.HandlerSuccess(response.HandlerSuccessParams{
		Context: context,
		Payload: &user,
		Status:  http.StatusOK,
	})
}

type handleLoginParams struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// todo - add support for email along with username
func handleLogin(context echo.Context) error {
	var params handleLoginParams
	var user User

	err := context.Bind(&params)
	if err != nil {
		return response.HandlerError(response.HandlerErrorParams{
			Context: context,
			Err:     err,
			ErrFunc: "handleLogin context.Bind",
			Status:  http.StatusBadRequest,
		})
	}

	token, err := handleLoginMacro(&user, params.Username, params.Password)
	if err != nil {
		return response.HandlerError(response.HandlerErrorParams{
			Context: context,
			Err:     err,
			ErrFunc: "handleLogin handleLoginMacro",
			Status:  http.StatusBadRequest,
		})
	}

	return response.HandlerSuccess(response.HandlerSuccessParams{
		Context: context,
		Payload: map[string]interface{}{
			"token": token,
		},
		Status: http.StatusOK,
	})
}

// func handleGetUsers(context echo.Context) error {
// 	var users []User
// 	actionGetUsers(&users)

// 	return context.JSON(http.StatusOK, users)
// }

// func handleGetUserByUsername(context echo.Context) error {
// 	username := context.Param(cc.USERNAME)

// 	foundUser, err := actionGetUserByUsername(username, cc.NO_PASSWORD, false)
// 	if err != nil {
// 		return response.HandlerError(response.HandlerErrorParams{
// 			Context: context,
// 			Err:     err,
// 			ErrFunc: "actionGetUserByUsername",
// 			Status:  http.StatusBadRequest,
// 		})
// 	}

// 	return response.HandlerSuccess(response.HandlerSuccessParams{
// 		Context: context,
// 		Payload: foundUser,
// 		Status:  http.StatusOK,
// 	})
// }

// func handleSignUp(context echo.Context) error {
// 	var userPropsPayload UserProps

// 	err := context.Bind(&userPropsPayload)
// 	if err != nil {
// 		errorString := "handleSignUp error: " + err.Error()
// 		return context.String(http.StatusBadRequest, errorString)
// 	}

// 	savedUser, err := actionCreateUser(&userPropsPayload)
// 	if err != nil {
// 		errorString := "actionCreateUser error: " + err.Error()
// 		return context.String(http.StatusBadRequest, errorString)
// 	}

// 	return context.JSON(http.StatusCreated, savedUser)
// }
