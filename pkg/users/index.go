package users

import (
	"github.com/labstack/echo/v4"
)

func AddUserActions(echoServer *echo.Echo) {
	echoServer.POST("/sign_up", handleSignUp)
	echoServer.POST("/log_in", handleLogin)

	echoServer.GET("/users", handleGetUsers)
	echoServer.GET("/users/username/:username", handleGetUserByUsername)
	// echoServer.GET("/users/:id", handleGetUserById)

	// echoServer.DELETE("/users/:id", handleDestroyUser)
	// echoServer.POST("/users", HandleCreateUser)
}
