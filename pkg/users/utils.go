package users

import (
	"github.com/frankmeza/roomchat/pkg/constants"
)

func createGetUserParams(params loginParams) GetUserParams {
	getUserParams := GetUserParams{
		Email:    params.Email,
		Username: params.Username,
	}

	if getUserParams.Email != "" {
		getUserParams.ParamName = constants.EMAIL
	}

	if getUserParams.Username != "" {
		getUserParams.ParamName = constants.USERNAME
	}

	return getUserParams
}

func getParamToUse(params GetUserParams) string {
	var paramToUse string

	switch paramName := params.ParamName; paramName {
	case constants.EMAIL:
		paramToUse = params.Email
	case constants.ID:
		paramToUse = params.ID
	case constants.USERNAME:
		paramToUse = params.Username
	case constants.UUID:
		paramToUse = params.Uuid
	default:
		paramToUse = ""
	}

	return paramToUse
}
