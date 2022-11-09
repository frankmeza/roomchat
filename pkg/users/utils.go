package users

import (
	"github.com/frankmeza/roomchat/pkg/constants"
)

func createGetUserParams(params handleLoginParams) GetUserParams {
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
	if params.ParamName == constants.EMAIL {
		return params.Email
	}

	if params.ParamName == constants.ID {
		return params.ID
	}

	if params.ParamName == constants.USERNAME {
		return params.Username
	}

	if params.ParamName == constants.UUID {
		return params.Uuid
	}

	return ""
}
