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
