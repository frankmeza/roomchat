package users

import (
	"net/http"

	cc "github.com/frankmeza/roomchat/pkg/constants"
	"github.com/frankmeza/roomchat/pkg/db"
	"github.com/labstack/echo/v4"
	jsonMap "github.com/mitchellh/mapstructure"
)

func handleGetUsers(c echo.Context) error {
	conn := db.GetDbConnection()
	users := []User{}

	result := conn.Find(&users)
	if result.Error != nil {
		return echo.NewHTTPError(
			http.StatusNotFound,
			result.Error,
		)
	}

	return c.JSON(http.StatusOK, users)
}

func handleGetUserById(c echo.Context) error {
	conn := db.GetDbConnection()
	id := c.Param(cc.ID)

	user := &User{}
	params := GetUserParams{ID: id, ParamName: cc.ID}

	err := getUserByParam(conn, user, params)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusNotFound,
			err.Error(),
		)
	}

	return c.JSON(http.StatusOK, user.UserSpecs)
}

func handleGetUserByEmail(c echo.Context) error {
	conn := db.GetDbConnection()
	email := c.Param(cc.EMAIL)

	user := &User{}
	params := GetUserParams{Email: email, ParamName: cc.EMAIL}

	err := getUserByParam(conn, user, params)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusNotFound,
			err.Error(),
		)
	}

	return c.JSON(http.StatusOK, user)
}

func handleCreateUser(c echo.Context) error {
	conn := db.GetDbConnection()

	userSpecs := UserProps{
		Bio:        c.Param(cc.BIO),
		CurrentPic: c.Param(cc.CURRENT_PIC),
		Email:      c.Param(cc.EMAIL),
		Location:   c.Param(cc.LOCATION),
		Name:       c.Param(cc.NAME),
		Pics:       c.Param(cc.PICS),
	}

	user := &User{}

	err := jsonMap.Decode(userSpecs, &user)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			err.Error(),
		)
	}

	err = createNewUser(conn, user, c)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			err.Error(),
		)
	}

	return c.JSON(http.StatusCreated, user)
}

func handleDestroyUser(c echo.Context) error {
	conn := db.GetDbConnection()
	id := c.Param(cc.ID)

	user := &User{}
	params := GetUserParams{ID: id, ParamName: cc.ID}

	err := getUserByParam(conn, user, params)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusNotFound,
			err.Error(),
		)
	}

	if err := deleteUser(conn, user, c); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
