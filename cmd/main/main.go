package main

import (
	"flag"
	"os"

	"github.com/frankmeza/roomchat/pkg/connections"
	"github.com/frankmeza/roomchat/pkg/constants"
	"github.com/frankmeza/roomchat/pkg/db"
	"github.com/frankmeza/roomchat/pkg/errata"
	"github.com/frankmeza/roomchat/pkg/users"
	"github.com/frankmeza/roomchat/pkg/users/auth"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

const HOST_AND_PORT string = "127.0.0.1:9990"
const START_DEV_MESSAGE string = "append -dev 1 to start locally "

func makeDbMigrations(dbConn *gorm.DB) error {
	return dbConn.AutoMigrate(
		&connections.ConnectionProps{},
		&connections.Message{},
		&users.UserSession{},
		&users.UserSessionProps{},
		&users.User{},
	)
}

func addPkgActions(server *echo.Echo, authorizedGroup *echo.Group) {
	users.AddAuthenticationActions(server)

	connections.AddConnectionActions(authorizedGroup)
}

// basically fossilized at this point
func main() {
	dbConn, err := db.GetDbConnection()
	if err != nil {
		panic(errata.CreateError(err, errata.ErrMessage{
			Text: "main GetDbConnection",
		}))
	}

	err = makeDbMigrations(dbConn)
	if err != nil {
		panic(errata.CreateError(err, errata.ErrMessage{
			Text: "main AutoMigrate",
		}))
	}

	envFlag := flag.Int("dev", 0, START_DEV_MESSAGE+HOST_AND_PORT)

	flag.Parse()
	isDev := *envFlag == 1

	echoServer := echo.New()

	if isDev {
		echoServer.Use(middleware.Logger())
		echoServer.Use(middleware.Recover())
	}

	authorizedGroup := echoServer.Group("/auth")

	config := middleware.JWTConfig{
		Claims:     &auth.JwtClaims{},
		SigningKey: []byte(os.Getenv(constants.SIGNING_KEY)),
	}

	authorizedGroup.Use(middleware.JWTWithConfig(config))
	addPkgActions(echoServer, authorizedGroup)

	echoServer.Logger.Fatal(echoServer.Start(HOST_AND_PORT))
}
