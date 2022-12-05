package connections

import (
	"errors"
	"net/http"

	"github.com/frankmeza/roomchat/pkg/response"
	"github.com/labstack/echo/v4"
)

func handleMakeConnection(context echo.Context) error {
	var params handleMakeConnectionParams

	err := context.Bind(&params)
	if err != nil {
		return response.SendResponse(context, response.Response{
			Error:  errors.New("handleMakeConnection Bind"),
			Status: http.StatusBadRequest,
		})
	}

	connection, err := handleMakeConnectionMacro(params)
	if err != nil {
		return response.SendResponse(context, response.Response{
			Error:  errors.New("handleMakeConnection handleMakeConnectionMacro"),
			Status: http.StatusBadRequest,
		})
	}

	return response.SendResponse(context, response.Response{
		Payload: map[string]interface{}{
			"connection": connection,
		},
		Status: http.StatusOK,
	})
}

func handleAddMessage(context echo.Context) error {
	var params handleAddMessageParams

	err := context.Bind(&params)
	if err != nil {
		return response.SendResponse(context, response.Response{
			Error:  errors.New("handleAddMessage Bind"),
			Status: http.StatusBadRequest,
		})
	}

	err = handleAddMessageMacro(params)
	if err != nil {
		return response.SendResponse(context, response.Response{
			Error:  errors.New("handleAddMessage handleAddMessageMacro"),
			Status: http.StatusBadRequest,
		})
	}

	return nil
}
