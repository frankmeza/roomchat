package auth

import (
	"net/http"

	"github.com/frankmeza/roomchat/pkg/users"
	"github.com/labstack/echo/v4"
	jsonMap "github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
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

func handleSignUp(c echo.Context) error {
	var userPropsPayload users.UserProps

	err := c.Bind(&userPropsPayload)
	if err != nil {
		errorString := "handleSignUp error: " + err.Error()
		return c.String(http.StatusBadRequest, errorString)
	}

	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(userPropsPayload.Password),
		HASH_COST,
	)

	if err != nil {
		errorString := "bcrypt.GenerateFromPassword error:" + err.Error()
		return c.String(http.StatusBadRequest, errorString)
	}

	userPropsPayload.Password = string(passwordHash)

	var user users.User
	err = jsonMap.Decode(userPropsPayload, &user.UserProps)

	if err != nil {
		errorString := "jsonMap.Decode error: " + err.Error()
		return c.String(http.StatusBadRequest, errorString)
	}

	savedUser, err := users.HandleCreateUser(c, &user)
	if err != nil {
		errorString := "users.HandleCreateUser error: " + err.Error()
		return c.String(http.StatusBadRequest, errorString)
	}

	return c.JSON(http.StatusCreated, savedUser)
}
