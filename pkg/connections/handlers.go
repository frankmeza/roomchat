package connections

import (
	"net/http"

	"github.com/frankmeza/roomchat/pkg/response"
	"github.com/labstack/echo/v4"
)

func handleMakeConnection(context echo.Context) error {
	var params handleMakeConnectionParams

	err := context.Bind(&params)
	if err != nil {
		return response.HandlerError(context, err, response.HandlerErrorParams{
			ErrMsg: "handleMakeConnection context.Bind",
			Status: http.StatusBadRequest,
		})
	}

	connection, err := handleMakeConnectionMacro(params)
	if err != nil {
		return response.HandlerError(context, err, response.HandlerErrorParams{
			ErrMsg: "handleMakeConnection handleMakeConnectionMacro",
			Status: http.StatusBadRequest,
		})
	}

	return response.HandlerSuccess(context, response.HandlerSuccessParams{
		Payload: connection,
		Status:  http.StatusOK,
	})
}

func handleAddMessage(context echo.Context) error {
	var params handleAddMessageParams

	err := context.Bind(&params)
	if err != nil {
		return response.HandlerError(context, err, response.HandlerErrorParams{
			ErrMsg: "handleAddMessage context.Bind",
			Status: http.StatusBadRequest,
		})
	}

	err = handleAddMessageMacro(params)
	if err != nil {
		return response.HandlerError(context, err, response.HandlerErrorParams{
			ErrMsg: "handleAddMessage handleAddMessageMacro",
			Status: http.StatusBadRequest,
		})
	}

	return nil
}
