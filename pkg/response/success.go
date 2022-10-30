package response

import "github.com/labstack/echo/v4"

type HandlerSuccessParams struct {
	Context echo.Context
	Payload interface{}
	Status  int
}

func HandlerSuccess(params HandlerSuccessParams) error {
	return params.Context.JSON(params.Status, params.Payload)
}
