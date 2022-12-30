package response

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Error   error
	Payload map[string]interface{}
	Status  int
}

func SendResponse(context echo.Context, params Response) error {
	if params.Error != nil {
		error := fmt.Sprintf("error caused by: %s", params.Error.Error())

		return context.JSON(params.Status, map[string]string{
			"error": error,
		})
	}

	return context.JSON(params.Status, params.Payload)
}
