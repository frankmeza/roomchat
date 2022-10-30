package main

import (
	"flag"

	"github.com/frankmeza/roomchat/pkg/constants"
	"github.com/frankmeza/roomchat/pkg/db"
	"github.com/frankmeza/roomchat/pkg/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const HOST_AND_PORT string = "127.0.0.1:9990"

func main() {
	dbConn := db.GetDbConnection()

	err := dbConn.AutoMigrate(&users.User{})
	if err != nil {
		panic(constants.DB_ERROR)
	}

	envMessage := "append -dev 1 to start locally on" + HOST_AND_PORT
	envFlag := flag.Int("dev", 0, envMessage)

	flag.Parse()
	isDev := *envFlag == 1

	server := echo.New()

	if isDev {
		server.Use(middleware.Logger())
		server.Use(middleware.Recover())
	}

	users.AddUserActions(server)

	server.Logger.Fatal(server.Start(HOST_AND_PORT))
}
