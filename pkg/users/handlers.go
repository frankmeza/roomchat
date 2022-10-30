package users

import (
	"net/http"

	"github.com/frankmeza/roomchat/pkg/response"
	"github.com/labstack/echo/v4"
)

func handleSignUp(context echo.Context) error {
	var user User
	var userProps UserProps

	err := context.Bind(&userProps)
	if err != nil {
		return response.HandlerError(response.HandlerErrorParams{
			Context: context,
			Err:     err,
			ErrFunc: "handleSignUp context.Bind",
			Status:  http.StatusBadRequest,
		})
	}

	err = handleSignUpMacro(&user, &userProps)
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
	Email    string `json:"email"`
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

	token, err := handleLoginMacro(&user, params)
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
			"user":  user,
		},
		Status: http.StatusOK,
	})
}
