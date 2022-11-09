package response

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type HandlerErrorParams struct {
	Context echo.Context
	ErrMsg  string
	Err     error
	Status  int
}

func HandlerError(context echo.Context, err error, params HandlerErrorParams) error {
	errorMessage := fmt.Sprintf(params.ErrMsg, "caused error:", err)
	return context.String(params.Status, errorMessage)
}
