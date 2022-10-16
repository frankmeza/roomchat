package users

import (
	"net/http"

	cc "github.com/frankmeza/roomchat/pkg/constants"
	"github.com/frankmeza/roomchat/pkg/db"
	"github.com/labstack/echo/v4"
)

func handleGetUserById(c echo.Context) error {
	conn := db.GetDbConnection()
	id := c.Param(cc.ID)

	user := &User{}
	user, err := getUserByParam(
		conn,
		user,
		cc.ID,
		GetUserParams{ID: id},
	)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func handleGetUserByEmail(c echo.Context) error {
	conn := db.GetDbConnection()
	email := c.Param(cc.EMAIL)

	user := &User{}
	user, err := getUserByParam(
		conn,
		user,
		cc.EMAIL,
		GetUserParams{Email: email},
	)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func handleCreateUser(c echo.Context) error {
	conn := db.GetDbConnection()
	newUser := &UserProps{}

	user, err := createNewUser(conn, newUser, c)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			err.Error(),
		)
	}

	return c.JSON(http.StatusCreated, user)
}
