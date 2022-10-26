package auth

import (
	"net/http"
	"os"

	cc "github.com/frankmeza/roomchat/pkg/constants"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	HASH_COST = 10
)

type jwtClaims struct {
	Name    string `json:"name"`
	UUID    string `json:"uuid"`
	IsAdmin bool   `json:"admin"`
	jwt.StandardClaims
}

func AddSignUpAction(echoServer *echo.Echo) {
	echoServer.POST("/sign_up", handleSignUp)
}

func AddAuthActions(echoServer *echo.Echo) {
	authRoutes := echoServer.Group("/auth/true")

	jwtConfig := middleware.JWTConfig{
		Claims:     &jwtClaims{},
		SigningKey: []byte(os.Getenv(cc.SIGNING_KEY)),
	}

	authRoutes.Use(middleware.JWTWithConfig(jwtConfig))

	authRoutes.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hella lit")
	})
}
