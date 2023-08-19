package users

import (
	"github.com/labstack/echo/v4"
)

func AddAuthenticationActions(echoServer *echo.Echo) {
	echoServer.POST("/sign_up", handleSignUp)
	echoServer.POST("/log_in", handleLogin)
}

func AddUserFetchActions(authorizedGroup *echo.Group) {
	authorizedGroup.GET("/users/username/:username", handleGetUser)
}
