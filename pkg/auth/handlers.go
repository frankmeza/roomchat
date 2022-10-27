package auth

import (
	"net/http"

	"github.com/frankmeza/roomchat/pkg/users"
	"github.com/labstack/echo/v4"
)

// func handleLogin(c echo.Context) error {
// 	username := c.Param(cc.USERNAME)
// 	password := c.Param(cc.PASSWORD)

// 	uuid := uuid.NewV4().String()

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims{
// 		IsAdmin: true,
// 		Name:    username + password,
// 		UUID:    uuid,
// 	})

// 	tokenAsString, err := token.SignedString(
// 		[]byte(os.Getenv(cc.JWT_SECRET)),
// 	)

// 	if err != nil {
// 		return err
// 	}

// 	return c.JSON(http.StatusOK, echo.Map{
// 		"token": tokenAsString,
// 	})
// }

func handleSignUp(context echo.Context) error {
	var userPropsPayload users.UserProps

	err := context.Bind(&userPropsPayload)
	if err != nil {
		errorString := "handleSignUp error: " + err.Error()
		return context.String(http.StatusBadRequest, errorString)
	}

	savedUser, err := users.HandleCreateUser(
		context,
		&userPropsPayload,
		generatePasswordString,
	)

	if err != nil {
		errorString := "users.HandleCreateUser error: " + err.Error()
		return context.String(http.StatusBadRequest, errorString)
	}

	return context.JSON(http.StatusCreated, savedUser)
}
