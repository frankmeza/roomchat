package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/frankmeza/roomchat/pkg/auth"
	"github.com/frankmeza/roomchat/pkg/connections"
	"github.com/frankmeza/roomchat/pkg/constants"
	"github.com/frankmeza/roomchat/pkg/db"
	"github.com/frankmeza/roomchat/pkg/errata"
	"github.com/frankmeza/roomchat/pkg/sessions"
	"github.com/frankmeza/roomchat/pkg/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

const HOST_AND_PORT string = "127.0.0.1:9990"
const START_DEV_MESSAGE string = "append -dev 1 to start locally "

func addPkgActions(server *echo.Echo, authorizedGroup *echo.Group) {
	users.AddAuthenticationActions(server)

	// below use /auth as prefix
	users.AddUserFetchActions(authorizedGroup)
	connections.AddConnectionActions(authorizedGroup)
}

func makeDbMigrations(dbConn *gorm.DB) error {
	return dbConn.AutoMigrate(
		// connections
		&connections.Connection{},
		&connections.Message{},

		// users
		&users.User{},
		&sessions.Session{},
	)
}

func main() {
	dbConn, err := db.GetDbConnection()
	if err != nil {
		panic(errata.CreateError(err, []string{
			"main GetDbConnection",
		}))
	}

	err = makeDbMigrations(dbConn)
	if err != nil {
		panic(errata.CreateError(err, []string{
			"main AutoMigrate",
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

	echoServer.GET("/", func(context echo.Context) error {
		return context.JSON(http.StatusOK, map[string]bool{
			"oh health yeah": true,
		})
	})

	addPkgActions(echoServer, authorizedGroup)
	echoServer.Logger.Fatal(echoServer.Start(HOST_AND_PORT))
}
