package connections

import "github.com/labstack/echo/v4"

func AddConnectionActions(authorizedGroup *echo.Group) {
	authorizedGroup.POST("/connections", handleMakeConnection)
	authorizedGroup.POST("/messages", handleAddMessage)
}
