package main

import (
	"flag"

	"github.com/frankmeza/roomchat/pkg/connections"
	"github.com/frankmeza/roomchat/pkg/db"
	"github.com/frankmeza/roomchat/pkg/errata"
	"github.com/frankmeza/roomchat/pkg/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

const HOST_AND_PORT string = "127.0.0.1:9990"

func makeDbMigrations(dbConn *gorm.DB) error {
	return dbConn.AutoMigrate(
		&connections.Connection{},
		&connections.Message{},
		&users.User{},
	)
}

func addPkgActions(server *echo.Echo) {
	users.AddUserActions(server)
	connections.AddConnectionActions(server)
}

// basically fossilized at this point
func main() {
	dbConn, err := db.GetDbConnection()
	if err != nil {
		panic(errata.CreateError(errata.ErrataParams{
			Err:    err,
			ErrMsg: "main db.GetDbConnection",
		}))
	}

	err = makeDbMigrations(dbConn)
	if err != nil {
		panic(errata.CreateError(errata.ErrataParams{
			Err:    err,
			ErrMsg: "main dbConn.AutoMigrate",
		}))
	}

	envMessage := "append -dev 1 to start locally " + HOST_AND_PORT
	envFlag := flag.Int("dev", 0, envMessage)

	flag.Parse()
	isDev := *envFlag == 1

	server := echo.New()

	if isDev {
		server.Use(middleware.Logger())
		server.Use(middleware.Recover())
	}

	addPkgActions(server)
	server.Logger.Fatal(server.Start(HOST_AND_PORT))
}
