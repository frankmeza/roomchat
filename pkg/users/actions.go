package users

import (
	"errors"

	"github.com/frankmeza/roomchat/pkg/auth"
	cc "github.com/frankmeza/roomchat/pkg/constants"
	"github.com/frankmeza/roomchat/pkg/db"
	"github.com/twinj/uuid"
)

func actionCreateUser(userPropsPayload *UserProps) (User, error) {
	var user User
	uuidString := uuid.NewV4().String()

	passwordHash, err := auth.GeneratePasswordString(userPropsPayload.Password)
	if err != nil {
		return User{}, err
	}

	err = useUsersAPI().CreateUser(
		&user,
		userPropsPayload,
		string(passwordHash),
		uuidString,
	)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func actionGetUsers(users *[]User) error {
	dbConn := db.GetDbConnection()

	err := getUsersDb(dbConn, users)
	if err != nil {
		return err
	}

	return nil
}

// func actionGetUserById(c echo.Context) error {
// 	dbConn := db.GetDbConnection()
// 	id := c.Param(cc.ID)

// 	user := User{}
// 	params := GetUserParams{ID: id, ParamName: cc.ID}

// 	err := getUserByParam(dbConn, &user, params)
// 	if err != nil {
// 		return c.String(
// 			http.StatusNotFound,
// 			err.Error(),
// 		)
// 	}

// 	return c.JSON(http.StatusOK, user)
// }

// func actionGetUserByEmail(c echo.Context) error {
// 	dbConn := db.GetDbConnection()
// 	email := c.Param(cc.EMAIL)

// 	user := User{}
// 	params := GetUserParams{Email: email, ParamName: cc.EMAIL}

// 	err := getUserByParam(dbConn, &user, params)
// 	if err != nil {
// 		return c.String(
// 			http.StatusNotFound,
// 			err.Error(),
// 		)
// 	}

// 	return c.JSON(http.StatusOK, user)
// }

func actionGetUserByUsername(username, password string, needsPassword bool) (
	User, error,
) {
	dbConn := db.GetDbConnection()

	user := User{}
	params := GetUserParams{Username: username, ParamName: cc.USERNAME}

	err := getUserByParam(dbConn, &user, params)
	if err != nil {
		return User{}, err
	}

	if !needsPassword {
		return user, nil
	}

	doesPasswordMatch := auth.CheckPasswordHash(user.UserProps.Password, password)
	if doesPasswordMatch {
		return user, nil
	}

	return User{}, errors.New("username or password is not correct")
}

// func actionDestroyUser(c echo.Context) error {
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
