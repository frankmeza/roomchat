package users

import (
	"github.com/labstack/echo/v4"
)

func AddUserActions(echoServer *echo.Echo) {
	echoServer.GET("/users", handleGetUsers)
	echoServer.GET("/users/:email", handleGetUserByEmail)
	echoServer.GET("/users/:id", handleGetUserById)

	// echoServer.DELETE("/users/:id", handleDestroyUser)
	// echoServer.POST("/users", HandleCreateUser)
}
