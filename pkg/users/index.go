package users

import (
	"github.com/labstack/echo/v4"
)

func AddUserActions(echoServer *echo.Echo) {
	echoServer.GET("/users/:email", handleGetUserByEmail)
	echoServer.GET("/users/:id", handleGetUserById)

	echoServer.POST("/users", handleCreateUser)
}
