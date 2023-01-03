package users

import (
	"errors"
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
		return response.SendResponse(context, response.Response{
			Error:  errors.New("handleGetUser GetUserByParam"),
			Status: http.StatusBadRequest,
		})
	}

	return response.SendResponse(context, response.Response{
		Payload: map[string]interface{}{
			"user": user,
		},
		Status: http.StatusOK,
	})
}

func handleSignUp(context echo.Context) error {
	var user User

	err := context.Bind(&user.UserProps)
	if err != nil {
		return response.SendResponse(context, response.Response{
			Error:  errors.New("handleSignUp context.Bind"),
			Status: http.StatusBadRequest,
		})
	}

	err = handleSignUpMacro(&user)
	if err != nil {
		return response.SendResponse(context, response.Response{
			Error:  errors.New("handleSignUp handleSignUpMacro"),
			Status: http.StatusBadRequest,
		})
	}

	return response.SendResponse(context, response.Response{
		Payload: map[string]interface{}{
			"user": user,
		},
		Status: http.StatusOK,
	})
}

func handleLogin(context echo.Context) error {
	var params loginParams
	var user User

	err := context.Bind(&params)
	if err != nil {
		return response.SendResponse(context, response.Response{
			Error:  errors.New("handleLogin context.Bind"),
			Status: http.StatusBadRequest,
		})
	}

	loginMetadata, err := handleLoginMacro(user, params)
	if err != nil {
		return response.SendResponse(context, response.Response{
			Error:  errors.New("handleLogin handleLoginMacro"),
			Status: http.StatusBadRequest,
		})
	}

	return response.SendResponse(context, response.Response{
		Payload: map[string]interface{}{
			"session": loginMetadata.session,
			"token":   loginMetadata.token,
			"user":    user,
		},
		Status: http.StatusOK,
	})
}

func handleUpdateUser(context echo.Context) error {
	return response.SendResponse(context, response.Response{
		Payload: map[string]interface{}{"user": "so lit"},
		Status:  http.StatusOK,
	})
}
