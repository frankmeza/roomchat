package response

import "github.com/labstack/echo/v4"

type HandlerSuccessParams struct {
	Payload interface{}
	Status  int
}

func HandlerSuccess(context echo.Context, params HandlerSuccessParams) error {
	return context.JSON(params.Status, params.Payload)
}
