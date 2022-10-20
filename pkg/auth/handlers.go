package auth

import (
	"net/http"
	"os"

	cc "github.com/frankmeza/roomchat/pkg/constants"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/twinj/uuid"
)

func handleLogin(c echo.Context) error {
	username := c.Param(cc.USERNAME)
	password := c.Param(cc.PASSWORD)

	uuid := uuid.NewV4().String()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims{
		IsAdmin: true,
		Name:    username + password,
		UUID:    uuid,
	})

	tokenAsString, err := token.SignedString(
		[]byte(os.Getenv(cc.JWT_SECRET)),
	)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": tokenAsString,
	})
}
