package response

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type HandlerErrorParams struct {
	Context echo.Context
	ErrFunc string
	Err     error
	Status  int
}

func HandlerError(params HandlerErrorParams) error {
	errorMessage := fmt.Sprintf(params.ErrFunc, "caused error:", params.Err)
	return params.Context.String(params.Status, errorMessage)
}
