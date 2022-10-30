package errata

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

// handlers return http.status errors, they call actions
// actions return action (object, errors); they call api, utils
// api return api error, mutate container object; they call db

type HandlerErrorParams struct {
	CallingFn string
	Context   echo.Context
	Err       error
	Status    int
}

func HandlerError(params HandlerErrorParams) error {
	errorMessage := fmt.Sprintf(params.CallingFn, "caused error:", params.Err)

	return params.Context.String(params.Status, errorMessage)
}
