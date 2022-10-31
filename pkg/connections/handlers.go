package connections

import (
	"net/http"

	"github.com/frankmeza/roomchat/pkg/response"
	"github.com/labstack/echo/v4"
)

func handleMakeConnection(context echo.Context) error {
	var connection Connection

	err := context.Bind(&connection)
	if err != nil {
		return response.HandlerError(response.HandlerErrorParams{
			Context: context,
			Err:     err,
			ErrFunc: "handleMakeConnection context.Bind",
			Status:  http.StatusBadRequest,
		})
	}

	err = handleMakeConnectionMacro(&connection)
	if err != nil {
		return response.HandlerError(response.HandlerErrorParams{
			Context: context,
			Err:     err,
			ErrFunc: "handleMakeConnection handleMakeConnectionMacro",
			Status:  http.StatusBadRequest,
		})
	}

	return response.HandlerSuccess(response.HandlerSuccessParams{
		Context: context,
		Payload: connection,
		Status:  http.StatusOK,
	})
}
