package db

import (
	"os"

	"github.com/frankmeza/roomchat/pkg/constants"
	cc "github.com/frankmeza/roomchat/pkg/constants"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDbConnection() *gorm.DB {
	postgresDb := postgres.Open(os.Getenv(cc.DB_URL))
	dbConn, err := gorm.Open(postgresDb, &gorm.Config{})

	if err != nil {
		panic(constants.DB_ERROR)
	}

	return dbConn
}
