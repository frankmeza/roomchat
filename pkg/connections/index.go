package connections

import "github.com/labstack/echo/v4"

func AddConnectionActions(echoServer *echo.Echo) {
	echoServer.POST("/connections", handleMakeConnection)
	echoServer.POST("/messages", handleAddMessage)
}
