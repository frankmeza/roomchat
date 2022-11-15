package users

import (
	"net/http"

	"github.com/frankmeza/roomchat/pkg/constants"
	"github.com/frankmeza/roomchat/pkg/response"
	"github.com/labstack/echo/v4"
)

func handleGetUser(context echo.Context) error {
	var user User

	params := GetUserParams{
		ParamName: constants.USERNAME,
		Username:  context.Param(constants.USERNAME),
	}

	err := UseUsersAPI().GetUserByParam(&user, params)
	if err != nil {
		return response.HandlerError(context, err, response.HandlerErrorParams{
			ErrMsg: "handleGetUser GetUserByParam",
			Status: http.StatusBadRequest,
		})
	}

	return response.HandlerSuccess(context, response.HandlerSuccessParams{
		Payload: &user,
		Status:  http.StatusOK,
	})
}

func handleSignUp(context echo.Context) error {
	var user User

	err := context.Bind(&user.UserProps)
	if err != nil {
		return response.HandlerError(context, err, response.HandlerErrorParams{
			ErrMsg: "handleSignUp context.Bind",
			Status: http.StatusBadRequest,
		})
	}

	err = handleSignUpMacro(&user)
	if err != nil {
		return response.HandlerError(context, err, response.HandlerErrorParams{
			ErrMsg: "handleSignUp handleSignUpMacro",
			Status: http.StatusBadRequest,
		})
	}

	return response.HandlerSuccess(context, response.HandlerSuccessParams{
		Payload: &user,
		Status:  http.StatusOK,
	})
}

func handleLogin(context echo.Context) error {
	var params handleLoginParams
	var user User

	err := context.Bind(&params)
	if err != nil {
		return response.HandlerError(context, err, response.HandlerErrorParams{
			ErrMsg: "handleLogin context.Bind",
			Status: http.StatusBadRequest,
		})
	}

	loginMetadata, err := handleLoginMacro(user, params)
	if err != nil {
		return response.HandlerError(context, err, response.HandlerErrorParams{
			ErrMsg: "handleLogin handleLoginMacro",
			Status: http.StatusBadRequest,
		})
	}

	return response.HandlerSuccess(context, response.HandlerSuccessParams{
		Payload: map[string]interface{}{
			"session": loginMetadata.session,
			"token":   loginMetadata.token,
			"user":    user,
		},
		Status: http.StatusOK,
	})
}

func handleUpdateUser(context echo.Context) error {
	return response.HandlerSuccess(context, response.HandlerSuccessParams{
		Payload: map[string]string{"user": "so lit"},
		Status:  http.StatusOK,
	})
}
