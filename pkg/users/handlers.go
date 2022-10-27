package users

import (
	"net/http"

	cc "github.com/frankmeza/roomchat/pkg/constants"
	"github.com/frankmeza/roomchat/pkg/db"
	"github.com/labstack/echo/v4"
	jsonMap "github.com/mitchellh/mapstructure"
	"github.com/twinj/uuid"
)

func handleGetUsers(c echo.Context) error {
	dbConn := db.GetDbConnection()
	users := []User{}

	result := dbConn.Find(&users)
	if result.Error != nil {
		return c.String(
			http.StatusNotFound,
			result.Error.Error(),
		)
	}

	return c.JSON(http.StatusOK, users)
}

func handleGetUserById(c echo.Context) error {
	dbConn := db.GetDbConnection()
	id := c.Param(cc.ID)

	user := &User{}
	params := GetUserParams{ID: id, ParamName: cc.ID}

	err := getUserByParam(dbConn, user, params)
	if err != nil {
		return c.String(
			http.StatusNotFound,
			err.Error(),
		)
	}

	return c.JSON(http.StatusOK, user)
}

func handleGetUserByEmail(c echo.Context) error {
	dbConn := db.GetDbConnection()
	email := c.Param(cc.EMAIL)

	user := &User{}
	params := GetUserParams{Email: email, ParamName: cc.EMAIL}

	err := getUserByParam(dbConn, user, params)
	if err != nil {
		return c.String(
			http.StatusNotFound,
			err.Error(),
		)
	}

	return c.JSON(http.StatusOK, user)
}

func HandleCreateUser(
	c echo.Context,
	userPropsPayload *UserProps,
	generatePasswordString func(plaintext string) (string, error),
) (User, error) {
	uuidString := uuid.NewV4().String()

	var user User
	user.Uuid = uuidString

	passwordHash, err := generatePasswordString(userPropsPayload.Password)
	if err != nil {
		return User{}, err
	}

	userPropsPayload.Password = string(passwordHash)
	userPropsPayload.Uuid = uuidString

	err = jsonMap.Decode(userPropsPayload, &user.UserProps)
	if err != nil {
		return User{}, err
	}

	err = saveUserDb(c, &user)
	if err != nil {
		return user, err
	}

	return user, nil
}

// func handleDestroyUser(c echo.Context) error {
// 	dbConn := db.GetDbConnection()
// 	id := c.Param(cc.ID)

// 	user := &User{}
// 	params := GetUserParams{ID: id, ParamName: cc.ID}

// 	err := getUserByParam(dbConn, user, params)
// 	if err != nil {
// 		return c.String(
// 			http.StatusNotFound,
// 			err.Error(),
// 		)
// 	}

// 	if err := deleteUser(dbConn, user, c); err != nil {
// 		return err
// 	}

// 	return c.NoContent(http.StatusOK)
// }
